package system

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
	"github.com/satori/go.uuid"
)

type User struct {
	global.MODEL
	UUID     uuid.UUID `json:"uuid"`     // 用户UUID
	Username string    `json:"userName"` // 用户登录名
	Password string    `json:"-"`        // 用户登录密码
	NickName string    `json:"nickName"` // 用户昵称
	Avatar   string    `json:"avatar"`   // 用户头像
	Status   int       `json:"status"`   // 状态，1、正常 2、禁用
}

type UserRoles struct {
	UserId snowflake.ID `json:"userId"` // 用户ID
	RoleId snowflake.ID `json:"roleId"` // 角色ID
}
