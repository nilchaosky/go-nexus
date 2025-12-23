package sqlite

import (
	"github.com/nilchaosky/go-nexus/serialize/variant"
)

// Snowflake 雪花ID主键结构体
type Snowflake struct {
	ID variant.SerializeInt64 `json:"id" gorm:"column:id;primaryKey;type:integer;comment:ID"`
}

// GetID 获取主键ID
func (s *Snowflake) GetID() int64 {
	return s.ID.Int64()
}

// SetID 设置主键ID
func (s *Snowflake) SetID(id int64) {
	s.ID = variant.NewSerializeInt64(id)
}

// AutoIncrement 自增ID主键结构体
type AutoIncrement struct {
	ID variant.SerializeInt64 `json:"id" gorm:"column:id;primaryKey;autoIncrement;type:integer;comment:ID"`
}

// GetID 获取主键ID
func (a *AutoIncrement) GetID() int64 {
	return a.ID.Int64()
}

// SetID 设置主键ID
func (a *AutoIncrement) SetID(id int64) {
	a.ID = variant.NewSerializeInt64(id)
}
