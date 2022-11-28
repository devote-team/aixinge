package global

import (
	"aixinge/utils/snowflake"
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID        snowflake.ID   `json:"id,omitempty" swaggertype:"string"` // 主键ID
	CreatedAt time.Time      `json:"createdAt,omitempty"`               // 创建时间
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`               // 更新时间
	DeletedAt gorm.DeletedAt `json:"-"`                                 // 删除时间
}
