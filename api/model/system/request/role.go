package request

import "aixinge/utils/snowflake"

// RoleMenuParams 角色分配菜单参数对象
type RoleMenuParams struct {
	ID      snowflake.ID   `json:"id,omitempty"` // 角色ID
	MenuIds []snowflake.ID `json:"menuIds"`      // 菜单ID集合
}

// RoleUserParams 角色分配用户参数对象
type RoleUserParams struct {
	ID      snowflake.ID   `json:"id,omitempty"` // 角色ID
	UserIds []snowflake.ID `json:"userIds"`      // 用户ID集合
}
