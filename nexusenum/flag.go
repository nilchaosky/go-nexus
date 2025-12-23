package nexusenum

// Flag 标志枚举
type Flag int

const (
	// FlagNo 否
	FlagNo Flag = 0
	// FlagYes 是
	FlagYes Flag = 1
)

// String 返回标志字符串
func (f Flag) String() string {
	switch f {
	case FlagNo:
		return "否"
	case FlagYes:
		return "是"
	default:
		return "未知"
	}
}

// Value 返回标志值
func (f Flag) Value() int {
	return int(f)
}

// Bool 转换为布尔值
func (f Flag) Bool() bool {
	return f == FlagYes
}

// FromBool 从布尔值创建标志
func FromBool(b bool) Flag {
	if b {
		return FlagYes
	}
	return FlagNo
}
