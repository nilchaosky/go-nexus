package redis

import (
	"context"
	"errors"
	"reflect"

	"github.com/nilchaosky/go-nexus/serialize"
	"github.com/nilchaosky/go-nexus/utils"
	"github.com/redis/go-redis/v9"
)

// Client 客户端包装结构体
// 实现了 Generic、String、List、Set、Hash、ZSet 接口
type Client struct {
	redis.UniversalClient
}

// NewClient 创建客户端
func NewClient(client redis.UniversalClient) *Client {
	return &Client{UniversalClient: client}
}

// GetRawClient 获取原始客户端
func (c *Client) GetRawClient() redis.UniversalClient {
	return c.UniversalClient
}

// Ping 测试连接
func (c *Client) Ping(ctx context.Context) error {
	return c.UniversalClient.Ping(ctx).Err()
}

// Close 关闭客户端连接
func (c *Client) Close() error {
	return c.UniversalClient.Close()
}

// unmarshalValue 获取键值并反序列化到结构体
func (c *Client) unmarshalValue(value interface{}, fn func() (string, error)) error {
	_, err := utils.IsPointer(value)
	if err != nil {
		return err
	}

	data, err := fn()
	if err != nil {
		return err
	}

	if data == "" {
		return errors.New("键值为空")
	}

	if err := serialize.JSONIter.Unmarshal([]byte(data), value); err != nil {
		return err
	}

	return nil
}

// unmarshalSlice 获取字符串切片并反序列化到结构体切片
func (c *Client) unmarshalSlice(value interface{}, fn func() ([]string, error)) error {
	rv, err := utils.IsSlice(value)
	if err != nil {
		return err
	}

	results, err := fn()
	if err != nil {
		return err
	}

	sliceType := rv.Type()
	elemType := sliceType.Elem()
	newSlice := reflect.MakeSlice(sliceType, 0, len(results))

	for _, result := range results {
		if result == "" {
			continue
		}

		elemValue := reflect.New(elemType).Interface()
		if err := serialize.JSONIter.Unmarshal([]byte(result), elemValue); err != nil {
			return err
		}

		elemPtr := reflect.ValueOf(elemValue)
		newSlice = reflect.Append(newSlice, elemPtr.Elem())
	}

	rv.Set(newSlice)
	return nil
}
