package nexusenum

// Method 方法枚举
type Method int

const (
	// MethodGet GET方法
	MethodGet Method = 1
	// MethodPost POST方法
	MethodPost Method = 2
	// MethodPut PUT方法
	MethodPut Method = 3
	// MethodDelete DELETE方法
	MethodDelete Method = 4
	// MethodPatch PATCH方法
	MethodPatch Method = 5
)

// String 返回方法字符串
func (m Method) String() string {
	switch m {
	case MethodGet:
		return "GET"
	case MethodPost:
		return "POST"
	case MethodPut:
		return "PUT"
	case MethodDelete:
		return "DELETE"
	case MethodPatch:
		return "PATCH"
	default:
		return "未知"
	}
}

// Value 返回方法值
func (m Method) Value() int {
	return int(m)
}
