package message

import (
	"aixinge/api/model/common"
	"aixinge/global"
	"aixinge/utils/snowflake"
	"database/sql/driver"
	"encoding/json"
)

type MailLog struct {
	global.MODEL
	AppId       snowflake.ID       `json:"appId,omitempty" swaggertype:"string"`      // 应用 ID
	TemplateId  snowflake.ID       `json:"templateId,omitempty" swaggertype:"string"` // 邮件模板 ID
	RequestId   snowflake.ID       `json:"requestId,omitempty" swaggertype:"string"`  // 唯一请求 ID
	To          MailAddress        `json:"to" gorm:"type:json"`                       // 发件地址集合
	Cc          MailAddress        `json:"cc" gorm:"type:json"`                       // 抄送地址集合
	Bcc         MailAddress        `json:"bcc" gorm:"type:json"`                      // 密送地址集合
	Parameters  string             `json:"parameters"`                                // 邮件参数
	Content     string             `json:"content"`                                   // 邮件具体内容
	Attachments common.Attachments `json:"attachments" gorm:"type:json"`              // 附件JSON
	Status      int                `json:"status"`                                    // 状态 1、正常 2、异常
	ErrMsg      string             `json:"errMsg"`                                    // 错误日志

}

type MailAddress []string

func (m MailAddress) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *MailAddress) Scan(src any) error {
	return json.Unmarshal(src.([]byte), m)
}
