package common

import (
	"aixinge/utils/snowflake"
	"database/sql/driver"
	"encoding/json"
)

type Attachments []Attachment

func (a Attachments) Value() (driver.Value, error) {
	b, err := json.Marshal(a)
	return string(b), err
}

func (a *Attachments) Scan(src any) error {
	return json.Unmarshal(src.([]byte), a)
}

type Attachment struct {
	FileId   snowflake.ID `json:"fileId" swaggertype:"string"` // 文件 ID
	FileName string       `json:"fileName"`                    // 文件名称
}
