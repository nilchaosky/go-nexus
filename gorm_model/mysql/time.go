package mysql

import (
	"github.com/nilchaosky/go-nexus/serialize/variant"
	"gorm.io/gorm"
)

// Timestamps 时间戳结构体
type Timestamps struct {
	CreatedAt variant.SerializeTime `json:"created_at" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt variant.SerializeTime `json:"updated_at" gorm:"column:updated_at;type:datetime;comment:更新时间"`
}

// SoftDelete 软删除结构体
type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;index;comment:删除时间"`
}
