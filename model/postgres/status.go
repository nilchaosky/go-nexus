package postgres

import (
	"github.com/nilchaosky/go-nexus/nexusenum"
)

// Status 状态结构体
type Status struct {
	Status nexusenum.Status `json:"status" gorm:"column:status;type:integer;default:1;comment:状态"`
}
