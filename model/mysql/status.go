package mysql

import (
	"github.com/nilchaosky/go-nexus/nexusenum"
)

// Status 状态结构体
type Status struct {
	Status nexusenum.Status `json:"status" gorm:"column:status;not null;default:1;comment:状态"`
}
