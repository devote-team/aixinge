package request

import (
	"aixinge/api/model/common/request"
	"aixinge/api/model/system"
	"aixinge/global"
)

// AddMenuAuthorityInfo Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.BaseMenu
	AuthorityId string // 角色ID
}

func DefaultMenu() []system.BaseMenu {
	return []system.BaseMenu{{
		MODEL:     global.MODEL{ID: 1},
		ParentId:  1,
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}

type MenuParams struct {
	Name string `json:"name"` // 菜单名称
}

type MenuPageParams struct {
	request.PageInfo
	Title  string `json:"title"`
	Status int    `json:"status,string,int"`
}
