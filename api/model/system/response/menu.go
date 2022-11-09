package response

import "aixinge/api/model/system"

type MenuResponse struct {
	Menu system.Menu `json:"menu"`
}

type MenuTreeResponse struct {
	system.Menu
	NoCache  int                 `json:"noCache"`            // 不缓存 1，是 2，否
	Title    string              `json:"title"`              // 菜单名
	Icon     string              `json:"icon"`               // 菜单图标
	Remark   string              `json:"remark"`             // 备注
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
