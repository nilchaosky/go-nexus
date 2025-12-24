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

// MarshalJSON 序列化为DateTime格式JSON字符串
func (t SerializeTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(t.Time.Format(time.DateTime))
}

// UnmarshalJSON 从JSON字符串反序列化
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

// String 返回DateTime格式字符串
func (t SerializeTime) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(time.DateTime)
}

// GetTime 返回time.Time值
func (t SerializeTime) GetTime() time.Time {
	return t.Time
}

// NewSerializeTime 创建SerializeTime实例
func NewSerializeTime(v time.Time) SerializeTime {
	return SerializeTime{Time: v}
}

// Value 转换为数据库驱动值
func (t SerializeTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// Scan 从数据库驱动值扫描
func (t *SerializeTime) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		return t.scanFromBytes(v)
	case string:
		return t.scanFromString(v)
	default:
		return fmt.Errorf("无法扫描 %T 到 SerializeTime", value)
	}
}

// scanFromBytes 从字节数组扫描时间
func (t *SerializeTime) scanFromBytes(data []byte) error {
	if len(data) == 0 {
		t.Time = time.Time{}
		return nil
	}

	str := string(data)
	return t.scanFromString(str)
}

// scanFromString 从字符串扫描时间
func (t *SerializeTime) scanFromString(str string) error {
	if str == "" {
		t.Time = time.Time{}
		return nil
	}

	// 尝试多种时间格式
	formats := []string{
		time.DateTime,
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
	}

	for _, format := range formats {
		if parsedTime, err := time.Parse(format, str); err == nil {
			t.Time = parsedTime
			return nil
		}
	}

	return fmt.Errorf("无法解析时间格式: %s", str)
}

// DateOnly 返回DateOnly格式字符串
func (t SerializeTime) DateOnly() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Time.Format(time.DateOnly)
}
