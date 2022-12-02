package message

import (
	"aixinge/utils/snowflake"
)

type ChannelTemplate struct {
	ChannelId  snowflake.ID `json:"channelId,omitempty" swaggertype:"string"`  // 渠道ID
	TemplateId snowflake.ID `json:"templateId,omitempty" swaggertype:"string"` // 模板ID
	Type       MsgType      `json:"type"`                                      // 消息类型(枚举)
	Default    int          `json:"default"`                                   // 默认渠道 1、是 2、否
}
