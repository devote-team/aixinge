package system

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
)

type BaseMenu struct {
	global.MODEL
	ParentId  snowflake.ID  `json:"parentId"`  // 父菜单ID
	Path      string        `json:"path" `     // 路由path
	Redirect  string        `json:"redirect"`  // 重定向的路由path
	Name      string        `json:"name"`      // 路由name
	Hidden    int           `json:"hidden"`    // 是否在列表隐藏 1，是 2，否
	Component string        `json:"component"` // 对应前端文件路径
	Sort      int           `json:"sort"`      // 排序标记
	IsFrame   int           `json:"isFrame"`   // Frame 1，是 2，否
	Meta      `json:"meta"` // 附加属性
}

type Menu struct {
	global.MODEL
	ParentId  snowflake.ID  `json:"parentId"`  // 父菜单ID
	Path      string        `json:"path" `     // 路由path
	Redirect  string        `json:"redirect"`  // 重定向的路由path
	Name      string        `json:"name"`      // 路由name
	Hidden    int           `json:"hidden"`    // 是否在列表隐藏 1，是 2，否
	Component string        `json:"component"` // 对应前端文件路径
	Sort      int           `json:"sort"`      // 排序标记
	IsFrame   int           `json:"isFrame"`   // Frame 1，是 2，否
	Status    int           `json:"status"`    // 状态，1、正常 2、禁用
	Meta      `json:"meta"` // 附加属性
}

type Meta struct {
	NoCache int    `json:"noCache"` // 不缓存 1，是 2，否
	Title   string `json:"title"`   // 菜单名
	Icon    string `json:"icon"`    // 菜单图标
	Remark  string `json:"remark"`  // 备注
}
