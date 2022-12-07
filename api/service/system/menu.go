package system

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/system"
	systemReq "aixinge/api/model/system/request"
	systemRes "aixinge/api/model/system/response"
	"aixinge/global"
	"aixinge/utils"
	"aixinge/utils/snowflake"
)

type MenuService struct {
}

func (c *MenuService) Create(menu system.Menu) error {
	menu.ID = utils.Id()
	// 状态，1、正常 2、禁用
	menu.Status = 1
	if menu.ParentId < 2 {
		// 设置父类为 1
		menu.ParentId = 1
	}
	return global.DB.Create(&menu).Error
}

func (c *MenuService) Delete(idsReq request.IdsReq) error {
	return global.DB.Delete(&[]system.Menu{}, "id in ?", idsReq.Ids).Error
}

func (c *MenuService) Update(menu system.Menu) (error, system.Menu) {
	if menu.ParentId < 2 {
		// 设置父类为 1
		menu.ParentId = 1
	}
	return global.DB.Updates(&menu).Error, menu
}

func (c *MenuService) GetById(id snowflake.ID) (err error, menu system.Menu) {
	err = global.DB.Where("id = ?", id).First(&menu).Error
	return err, menu
}

func (c *MenuService) Page(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.Menu{})
	var menuList []system.Menu
	err = db.Count(&total).Error
	err = db.Order("sort DESC").Limit(limit).Offset(offset).Find(&menuList).Error
	return err, menuList, total
}

func (c *MenuService) List(params systemReq.MenuParams) (err error, list interface{}) {
	db := global.DB.Model(&system.Menu{})
	if len(params.Name) > 0 {
		db.Where("name LIKE ?", params.Name)
	}
	var menuList []system.Menu
	err = db.Select("id", "parent_id", "name").Order("sort").Find(&menuList).Error
	return err, menuList
}

// AuthList 当前登录用户权限菜单
func (c *MenuService) AuthList(customClaims *systemReq.TokenClaims) (err error, list interface{}) {
	db := global.DB.Model(&system.Menu{})
	var menuList []system.Menu
	db.Raw("select m.* from axg_menu m join axg_role_menu r on m.id=r.menu_id join axg_user_role u on r.role_id=u.role_id where m.status=1 and u.user_id=?",
		customClaims.ID).Scan(&menuList)
	return err, menuList
}

func (c *MenuService) ListTree(info systemReq.MenuPageParams) (error, []*systemRes.MenuTreeResponse) {
	db := global.DB.Model(&system.Menu{})
	if info.Title != "" {
		db.Where("title like ?", "%"+info.Title+"%")
	}

	if info.Status != 0 {
		db.Where("status = ?", info.Status)
	}
	var menuList []system.Menu
	_ = db.Order("sort ASC").Find(&menuList).Error
	return GetMenuTree(menuList, 0)
}

// GetMenuTree 递归获取树形菜单
func GetMenuTree(menuList []system.Menu, parentId snowflake.ID) (error, []*systemRes.MenuTreeResponse) {
	tree := make([]*systemRes.MenuTreeResponse, 0)
	if len(menuList) > 0 {
		for _, item := range menuList {
			if item.ParentId == parentId {
				_, child := GetMenuTree(menuList, item.ID)
				node := &systemRes.MenuTreeResponse{
					Menu:     item,
					Children: child,
				}
				tree = append(tree, node)
			}
		}
	}
	return nil, tree
}
