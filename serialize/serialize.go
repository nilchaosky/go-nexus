package serialize

import (
	"github.com/nilchaosky/go-nexus/serialize/json"
	"github.com/nilchaosky/go-nexus/serialize/jsoniter"
	"github.com/nilchaosky/go-nexus/serialize/protobuf"
)

// Serializer 序列化接口
// 提供序列化和反序列化功能
type Serializer interface {
	// Marshal 将数据序列化为字节数组
	// v 为要序列化的数据
	// 返回序列化后的字节数组和错误
	Marshal(v interface{}) ([]byte, error)

	// Unmarshal 将字节数组反序列化为数据
	// data 为要反序列化的字节数组
	// v 为反序列化后的数据指针
	// 返回错误
	Unmarshal(data []byte, v interface{}) error
}

var (
	// JSON JSON序列化器（使用标准库encoding/json）
	JSON Serializer = json.NewSerializer()

	// JSONIter JSON序列化器（使用json-iterator库，性能更好）
	JSONIter Serializer = jsoniter.NewSerializer()

	// Protobuf Protobuf序列化器
	Protobuf Serializer = protobuf.NewSerializer()
)
