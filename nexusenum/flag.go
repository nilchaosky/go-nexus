package nexusenum

// Flag 标志枚举
type Flag int32

const (
	_flag Flag = iota
	// FlagYes 是
	FlagYes
	// FlagNo 否
	FlagNo
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
func (f Flag) Value() int32 {
	return int32(f)
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
