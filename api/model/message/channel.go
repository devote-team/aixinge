package message

import (
	"aixinge/global"
	"database/sql/driver"
	"encoding/json"
)

type ChannelConfig struct {
	Name string `json:"name"` // 渠道名称
}

func (c ChannelConfig) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ChannelConfig) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}

type Channel struct {
	global.MODEL
	Name     string          `json:"name"`     // 消息渠道名称
	Type     MsgType         `json:"type"`     // 消息类型(枚举)
	Provider ChannelProvider `json:"provider"` // 消息服务提供商
	Weight   int             `json:"weight"`   // 权重
	Config   ChannelConfig   `json:"config"`   // 消息渠道配置（JSON）
	Remark   string          `json:"remark"`   // 备注
	Status   int             `json:"status"`   // 状态 1、正常 2、禁用
}
