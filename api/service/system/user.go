package system

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/system"
	systemReq "aixinge/api/model/system/request"
	"aixinge/global"
	"aixinge/utils"
	"aixinge/utils/snowflake"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (b *UserService) Create(uc systemReq.UserCreate) (err error, userInter system.User) {
	var user system.User
	if !errors.Is(global.DB.Where("username = ?", uc.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		// 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	var u system.User
	u.ID = utils.Id()
	u.UUID = uuid.NewV4()
	u.Username = uc.Username
	u.Password = utils.GetByteMd5([]byte(uc.Password + u.UUID.String()))
	u.NickName = uc.NickName
	u.Status = 1
	err = global.DB.Create(&u).Error
	return err, u
}

func (b *UserService) Delete(idsReq request.IdsReq) (err error) {
	err = global.DB.Delete(&[]system.User{}, "id in ?", idsReq.Ids).Error
	return err
}

func (b *UserService) Update(reqUser system.User) (err error, user system.User) {
	err = global.DB.Updates(&reqUser).Error
	return err, reqUser
}

func (b *UserService) Login(u *system.User) (err error, userInter *system.User) {
	var user system.User
	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	pwd := utils.GetByteMd5([]byte(u.Password + user.UUID.String()))
	if err != nil || user.Password != pwd {
		return errors.New("用户密码错误"), userInter
	}
	return err, &user
}

func (b *UserService) ChangePassword(u *system.User, newPassword string) (err error, userInter *system.User) {
	var user system.User
	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil || user.Password != utils.GetByteMd5([]byte(u.Password+user.UUID.String())) {
		return errors.New("用户密码错误"), userInter
	}
	// 重置新密码
	err = global.DB.Model(&system.User{}).Where("id = ?", user.ID).Update("password", utils.GetByteMd5([]byte(newPassword+user.UUID.String()))).Error
	return err, u
}

func (b *UserService) AssignRole(params systemReq.UserRoleParams) (err error) {
	if params.ID < 1 {
		return errors.New("用户ID不能为空")
	}
	if len(params.RoleIds) == 0 {
		return errors.New("角色ID集合不能为空")
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&system.UserRoles{})
		err = db.Delete("user_id = ?", params.ID).Error
		if err != nil {
			return errors.New("分配角色历史数据删除失败")
		}
		var userRoles []system.UserRoles
		for i := range params.RoleIds {
			var ur system.UserRoles
			ur.UserId = params.ID
			ur.RoleId = params.RoleIds[i]
			userRoles = append(userRoles, ur)
		}
		err = db.CreateInBatches(&userRoles, 100).Error
		if err != nil {
			return errors.New("分配角色保存失败")
		}
		return nil
	})
}

func (b *UserService) SelectedRoles(id snowflake.ID) (err error, list interface{}) {
	var roleIds []snowflake.ID
	var userRoleList []system.UserRoles
	err = global.DB.Where("user_id=?", id).Find(&userRoleList).Error
	if len(userRoleList) > 0 {
		for i := range userRoleList {
			roleIds = append(roleIds, userRoleList[i].RoleId)
		}
	}
	return err, roleIds
}

func (b *UserService) GetByUuid(uuid string) (err error, user system.User) {
	err = global.DB.Where("uuid = ?", uuid).First(&user).Error
	return err, user
}

func (b *UserService) GetById(id snowflake.ID) (err error, user system.User) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return err, user
}

func (b *UserService) Page(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.User{})
	var userList []system.User
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return err, userList, total
}

func (b *UserService) List() (err error, list interface{}) {
	db := global.DB.Model(&system.User{})
	var userList []system.User
	err = db.Where("status=?", 1).Find(&userList).Error
	return err, userList
}
