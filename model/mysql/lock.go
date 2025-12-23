package mysql

import (
	"gorm.io/plugin/optimisticlock"
)

// OptimisticLock 乐观锁结构体
type OptimisticLock struct {
	Version optimisticlock.Version `json:"-" gorm:"column:_version;type:int;default:0;comment:版本号"`
}
