package system

import (
	"aixinge/global"
	"aixinge/utils/snowflake"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid"`     // 用户UUID
	Username string    `json:"username"` // 用户登录名
	Password string    `json:"-"`        // 用户登录密码
	Nickname string    `json:"nickname"` // 用户昵称
	Avatar   string    `json:"avatar"`   // 用户头像˚
	Status   int       `json:"status"`   // 状态，1、正常 2、禁用
}

type UserRole struct {
	UserId snowflake.ID `json:"userId" swaggertype:"string"` // 用户ID
	RoleId snowflake.ID `json:"roleId" swaggertype:"string"` // 角色ID
}
