package nexusenum

// Status 状态枚举
type Status int

const (
	// StatusDisabled 禁用状态
	StatusDisabled Status = 0
	// StatusEnabled 启用状态
	StatusEnabled Status = 1
)

// String 返回状态字符串
func (s Status) String() string {
	switch s {
	case StatusDisabled:
		return "禁用"
	case StatusEnabled:
		return "启用"
	default:
		return "未知"
	}
}

// Value 返回状态值
func (s Status) Value() int {
	return int(s)
}
