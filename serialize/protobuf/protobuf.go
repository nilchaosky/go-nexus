package protobuf

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

// Serializer Protobuf序列化器
// 实现serialize.Serializer接口，使用google.golang.org/protobuf/proto进行序列化和反序列化
type Serializer struct{}

// Marshal 将数据序列化为Protobuf字节数组
// v 为要序列化的数据，必须是proto.Message类型
// 返回Protobuf格式的字节数组和错误
func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, errors.New("数据必须是proto.Message类型")
	}
	return proto.Marshal(msg)
}

// Unmarshal 将Protobuf字节数组反序列化为数据
// data 为要反序列化的Protobuf字节数组
// v 为反序列化后的数据指针，必须是proto.Message类型
// 返回错误
func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	msg, ok := v.(proto.Message)
	if !ok {
		return errors.New("数据必须是proto.Message类型")
	}
	return proto.Unmarshal(data, msg)
}

// NewSerializer 创建新的Protobuf序列化器实例
// 返回Serializer实例
func NewSerializer() *Serializer {
	return &Serializer{}
}
