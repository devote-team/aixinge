package message

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
	"database/sql/driver"
	"encoding/json"
)

type attachmentIds []snowflake.ID

func (c attachmentIds) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *attachmentIds) Scan(src any) error {
	return json.Unmarshal(src.([]byte), c)
}

type MailTemplate struct {
	global.MODEL
	AppId       snowflake.ID  `json:"appId,omitempty" swaggertype:"string"` // 应用 ID
	Name        string        `json:"name"`                                 // 模板名称
	Content     string        `json:"content"`                              // 模板内容
	Type        int           `json:"type"`                                 // 模板类型（1-文本、2-HTML）
	Attachments attachmentIds `json:"attachments" gorm:"type:json"`         // 附件ID集合
	Remark      string        `json:"remark"`                               // 备注
	Status      int           `json:"status"`                               // 状态 1、正常 2、禁用
}
