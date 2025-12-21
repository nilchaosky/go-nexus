package variant

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// SerializeTime time.Time变体类型
// 支持Gin框架的JSON序列化，将time.Time序列化为DateTime格式字符串
// 实现json.Marshaler和json.Unmarshaler接口
type SerializeTime struct {
	time.Time
}

// MarshalJSON 实现json.Marshaler接口
// 将SerializeTime序列化为DateTime格式的JSON字符串
func (t SerializeTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(t.Time.Format(time.DateTime))
}

// UnmarshalJSON 实现json.Unmarshaler接口
// 从JSON字符串反序列化为SerializeTime
func (t *SerializeTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("无法解析SerializeTime字符串: %w", err)
	}

	if str == "" {
		t.Time = time.Time{}
		return nil
	}

	parsedTime, err := time.Parse(time.DateTime, str)
	if err != nil {
		return fmt.Errorf("无法解析时间格式: %w", err)
	}

	t.Time = parsedTime
	return nil
}

// String 实现fmt.Stringer接口
// 返回SerializeTime的DateTime格式字符串表示
func (t SerializeTime) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(time.DateTime)
}

// GetTime 返回time.Time类型的值
// 返回SerializeTime的time.Time值
func (t SerializeTime) GetTime() time.Time {
	return t.Time
}

// NewSerializeTime 创建新的SerializeTime变体
// v 为time.Time值
// 返回SerializeTime实例
func NewSerializeTime(v time.Time) SerializeTime {
	return SerializeTime{Time: v}
}

// Value 实现driver.Valuer接口
// 将SerializeTime转换为数据库驱动可以使用的值
// 返回time.Time类型的值，用于数据库存储
func (t SerializeTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// DateOnly 返回DateOnly格式的字符串
// 返回格式为"2006-01-02"的日期字符串，只包含日期部分
func (t SerializeTime) DateOnly() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(time.DateOnly)
}
