package response

import "aixinge/api/model/system"

type MenuResponse struct {
	Menu system.Menu `json:"menu"`
}

type MenuTreeResponse struct {
	system.Menu
	Children []*MenuTreeResponse `json:"children,omitempty"` // 子类
}

type SysMenusResponse struct {
	Menus []system.Menu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system.BaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system.BaseMenu `json:"menu"`
}
