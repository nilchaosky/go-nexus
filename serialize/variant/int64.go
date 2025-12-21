package variant

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// SerializeInt64 int64变体类型
// 支持Gin框架的JSON序列化，将int64序列化为字符串以避免JavaScript精度丢失
// 实现json.Marshaler和json.Unmarshaler接口
type SerializeInt64 int64

// MarshalJSON 序列化为JSON字符串
func (i SerializeInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(i), 10))
}

// UnmarshalJSON 从JSON字符串或数字反序列化
func (i *SerializeInt64) UnmarshalJSON(data []byte) error {
	// 尝试先解析为字符串
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("无法解析SerializeInt64字符串: %w", err)
		}
		*i = SerializeInt64(val)
		return nil
	}

	// 如果字符串解析失败，尝试解析为数字
	var num int64
	if err := json.Unmarshal(data, &num); err != nil {
		return fmt.Errorf("无法解析SerializeInt64: %w", err)
	}
	*i = SerializeInt64(num)
	return nil
}

// String 返回字符串表示
func (i SerializeInt64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Int64 返回int64值
func (i SerializeInt64) Int64() int64 {
	return int64(i)
}

// NewSerializeInt64 创建SerializeInt64实例
func NewSerializeInt64(v int64) SerializeInt64 {
	return SerializeInt64(v)
}

// Value 转换为数据库驱动值
func (i SerializeInt64) Value() (driver.Value, error) {
	return int64(i), nil
}

// ToInt64Slice 转换为[]int64
func ToInt64Slice(slice []SerializeInt64) []int64 {
	if slice == nil {
		return nil
	}
	result := make([]int64, len(slice))
	for i, v := range slice {
		result[i] = int64(v)
	}
	return result
}

// FromInt64Slice 转换为[]SerializeInt64
func FromInt64Slice(slice []int64) []SerializeInt64 {
	if slice == nil {
		return nil
	}
	result := make([]SerializeInt64, len(slice))
	for i, v := range slice {
		result[i] = SerializeInt64(v)
	}
	return result
}
