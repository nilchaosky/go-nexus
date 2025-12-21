package jsoniter

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Serializer JSON序列化器（使用json-iterator库）
// 实现serialize.Serializer接口，使用github.com/json-iterator/go进行序列化和反序列化
type Serializer struct{}

// Marshal 将数据序列化为JSON字节数组
// v 为要序列化的数据
// 返回JSON格式的字节数组和错误
func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal 将JSON字节数组反序列化为数据
// data 为要反序列化的JSON字节数组
// v 为反序列化后的数据指针
// 返回错误
func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// NewSerializer 创建新的JSON序列化器实例（使用json-iterator库）
// 返回Serializer实例
func NewSerializer() *Serializer {
	return &Serializer{}
}
