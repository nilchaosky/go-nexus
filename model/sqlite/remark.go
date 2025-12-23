package sqlite

// Remark 备注结构体
type Remark struct {
	Remark string `json:"remark" gorm:"column:remark;type:text;comment:备注"`
}
