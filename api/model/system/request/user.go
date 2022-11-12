package request

import "aixinge/utils/snowflake"

type UserCreate struct {
	Username string `json:"userName"` // 用户登录名
	Password string `json:"password"` // 用户登录密码
	NickName string `json:"nickName"` // 用户昵称
}

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"` // 刷新票据
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// UserRoleParams 用户分配角色参数对象
type UserRoleParams struct {
	ID      snowflake.ID   `json:"id,omitempty"` // 用户ID
	RoleIds []snowflake.ID `json:"roleIds"`      // 角色ID集合
}
