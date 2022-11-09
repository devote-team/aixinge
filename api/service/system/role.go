package system

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/system"
	systemReq "aixinge/api/model/system/request"
	"aixinge/global"
	"aixinge/utils"
	"aixinge/utils/snowflake"
	"errors"
	"gorm.io/gorm"
)

type RoleService struct{}

func (t *RoleService) Create(role system.Role) (err error) {
	role.ID = utils.Id()
	role.Status = 1
	err = global.DB.Create(&role).Error
	return err
}

func (t *RoleService) Delete(idsReq request.IdsReq) (err error) {
	err = global.DB.Delete(&[]system.Role{}, "id in ?", idsReq.Ids).Error
	return err
}

func (t *RoleService) Update(reqRole system.Role) (err error, role system.Role) {
	err = global.DB.Updates(&reqRole).Error
	return err, reqRole
}

func (t *RoleService) AssignUser(params systemReq.RoleUserParams) (err error) {
	if params.ID < 1 {
		return errors.New("角色ID不能为空")
	}
	if len(params.UserIds) == 0 {
		return errors.New("用户ID集合不能为空")
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&system.UserRoles{})
		err = db.Where("role_id = ?", params.ID).Delete(&system.UserRoles{}).Error
		if err != nil {
			return errors.New("分配用户历史数据删除失败")
		}
		var userRoles []system.UserRoles
		for i := range params.UserIds {
			var rm system.UserRoles
			rm.RoleId = params.ID
			rm.UserId = params.UserIds[i]
			userRoles = append(userRoles, rm)
		}
		err = db.CreateInBatches(&userRoles, 100).Error
		if err != nil {
			return errors.New("分配用户保存失败")
		}
		return nil
	})
}

func (t *RoleService) SelectedUsers(id snowflake.ID) (err error, list interface{}) {
	var userIds []snowflake.ID
	var userRoleList []system.UserRoles
	err = global.DB.Where("role_id=?", id).Find(&userRoleList).Error
	if len(userRoleList) > 0 {
		for i := range userRoleList {
			userIds = append(userIds, userRoleList[i].UserId)
		}
	}
	return err, userIds
}

func (t *RoleService) AssignMenu(params systemReq.RoleMenuParams) (err error) {
	if params.ID < 1 {
		return errors.New("角色ID不能为空")
	}
	if len(params.MenuIds) == 0 {
		return errors.New("菜单ID集合不能为空")
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Model(&system.RoleMenus{})
		err = db.Where("role_id = ?", params.ID).Delete(&system.RoleMenus{}).Error
		if err != nil {
			return errors.New("分配菜单历史数据删除失败")
		}
		var roleMenus []system.RoleMenus
		for i := range params.MenuIds {
			var rm system.RoleMenus
			rm.RoleId = params.ID
			rm.MenuId = params.MenuIds[i]
			roleMenus = append(roleMenus, rm)
		}
		err = db.CreateInBatches(&roleMenus, 100).Error
		if err != nil {
			return errors.New("分配菜单保存失败")
		}
		return nil
	})
}

func (t *RoleService) SelectedMenus(id snowflake.ID) (err error, list interface{}) {
	var menuIds []snowflake.ID
	var roleMenuList []system.RoleMenus
	err = global.DB.Where("role_id=?", id).Find(&roleMenuList).Error
	if len(roleMenuList) > 0 {
		for i := range roleMenuList {
			menuIds = append(menuIds, roleMenuList[i].MenuId)
		}
	}
	return err, menuIds
}

func (t *RoleService) GetById(id snowflake.ID) (err error, role system.Role) {
	err = global.DB.Where("id = ?", id).First(&role).Error
	return err, role
}

func (t *RoleService) Page(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.Role{})
	var roleList []system.Role
	err = db.Count(&total).Error
	err = db.Order("sort DESC").Limit(limit).Offset(offset).Find(&roleList).Error
	return err, roleList, total
}

func (t *RoleService) List() (err error, list interface{}) {
	db := global.DB.Model(&system.Role{})
	var roleList []system.Role
	err = db.Where("status=?", 1).Order("sort DESC").Find(&roleList).Error
	return err, roleList
}
