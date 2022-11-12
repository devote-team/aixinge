package message

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
)

// BaseTemplate 基础消息模板
type BaseTemplate struct {
	global.MODEL
	AppId   snowflake.ID `json:"appId"`   // 应用 ID
	Name    string       `json:"name"`    // 模板名称
	Type    int          `json:"type"`    // 消息类型
	Content string       `json:"content"` // 模板内容
	Status  string       `json:"status"`  // 状态
	Remark  string       `json:"remark"`  // 备注
}

// BaseChannel 基础消息渠道
type BaseChannel struct {
	global.MODEL
	AppId    snowflake.ID `json:"appId"`    // 应用 ID
	Name     string       `json:"name"`     // 渠道名称
	Type     int          `json:"type"`     // 消息类型
	Provider int          `json:"provider"` // 服务提供商
	Status   string       `json:"status"`   // 状态
	Remark   string       `json:"remark"`   // 备注
}
