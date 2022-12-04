package system

import (
	"aixinge/utils/snowflake"
	"time"
)

type File struct {
	ID          snowflake.ID `json:"id,omitempty" gorm:"primarykey" swaggertype:"string"`
	CreatedAt   time.Time    `json:"createdAt,omitempty"`
	Md5         string       `json:"md5"`
	Path        string       `json:"path"`
	Ext         string       `json:"ext"`
	ContentType string       `json:"contentType"`
	Size        int64        `json:"size"`
	Filename    string       `json:"filename"`
}
