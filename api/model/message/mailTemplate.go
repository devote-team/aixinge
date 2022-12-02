package message

import (
	"aixinge/api/model/common"
	"aixinge/global"
	"aixinge/utils/snowflake"
)

type MailTemplate struct {
	global.MODEL
	AppId       snowflake.ID       `json:"appId,omitempty" swaggertype:"string"` // 应用 ID
	Name        string             `json:"name"`                                 // 模板名称
	Content     string             `json:"content"`                              // 模板内容
	Type        int                `json:"type"`                                 // 模板类型（1-文本、2-HTML）
	Attachments common.Attachments `json:"attachments" gorm:"type:json"`         // 附件JSON
	Remark      string             `json:"remark"`                               // 备注
	Status      int                `json:"status"`                               // 状态 1、正常 2、禁用
}
