package mysql

// Remark 备注结构体
type Remark struct {
	Remark string `json:"remark" gorm:"column:remark;type:varchar(200);comment:备注"`
}
