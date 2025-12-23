package nexusenum

// Method 方法枚举
type Method string

const (
	// MethodGet GET方法
	MethodGet Method = "GET"
	// MethodPost POST方法
	MethodPost Method = "POST"
	// MethodPut PUT方法
	MethodPut Method = "PUT"
	// MethodDelete DELETE方法
	MethodDelete Method = "DELETE"
	// MethodPatch PATCH方法
	MethodPatch Method = "PATCH"
)

// String 返回方法字符串
func (m Method) String() string {
	return string(m)
}

// Value 返回方法值
func (m Method) Value() string {
	return string(m)
}
