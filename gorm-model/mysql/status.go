package mysql

import (
	"github.com/nilchaosky/go-nexus/nexusenum"
)

// Status 状态结构体
type Status struct {
	Value nexus_enum.Status `json:"status" gorm:"column:status;not null;default:1;comment:状态"`
}
