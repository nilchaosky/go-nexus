package redis

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/nilchaosky/go-nexus/nexusutils"
	"github.com/nilchaosky/go-nexus/serialize"
)

// String 字符串操作接口
type String interface {
	Cache(ctx context.Context, key string, value interface{}, expiration time.Duration, fn func() (interface{}, error)) error
	Get(ctx context.Context, key string) (string, error)
	GetStruct(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}) (bool, error)
	SetNXEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	SetXX(ctx context.Context, key string, value interface{}) (bool, error)
	SetXXEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	Append(ctx context.Context, key, value string) (int64, error)
	StrLen(ctx context.Context, key string) (int64, error)
	GetRange(ctx context.Context, key string, start, end int64) (string, error)
	SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	IncrByFloat(ctx context.Context, key string, value float64) (float64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, value int64) (int64, error)
	GetSet(ctx context.Context, key string, value interface{}) (string, error)
	GetSetStruct(ctx context.Context, key string, value interface{}, result interface{}) error
	MGet(ctx context.Context, keys ...string) ([]interface{}, error)
	MSet(ctx context.Context, values ...interface{}) error
	MSetNX(ctx context.Context, values ...interface{}) (bool, error)
}

// Cache 缓存方法：先从缓存获取，如果不存在则执行fn函数获取数据并缓存
func (c *Client) Cache(ctx context.Context, key string, value interface{}, expiration time.Duration, fn func() (interface{}, error)) error {
	rv, err := nexusutils.IsPointer(value)
	if err != nil {
		return err
	}
	rv = rv.Elem()

	// 先从缓存获取
	data, err := c.Get(ctx, key)
	if err == nil && data != "" {
		// 缓存存在，反序列化到结构体
		if err := serialize.JSONIter.Unmarshal([]byte(data), value); err != nil {
			return err
		}
		return nil
	}

	// 缓存不存在，执行fn函数获取数据
	result, err := fn()
	if err != nil {
		return err
	}

	// 直接赋值给value

	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() == reflect.Ptr {
		resultValue = resultValue.Elem()
	}

	if !resultValue.Type().AssignableTo(rv.Type()) {
		return errors.New("类型不匹配")
	}

	// 类型匹配，直接赋值
	rv.Set(resultValue)

	// 序列化并缓存数据
	jsonData, err := serialize.JSONIter.Marshal(result)
	if err != nil {
		return err
	}

	// 缓存数据
	if err := c.SetEX(ctx, key, string(jsonData), expiration); err != nil {
		return err
	}

	return nil
}

// Get 获取键值
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	result := c.UniversalClient.Get(ctx, key)
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Result()
}

// GetStruct 获取键值并反序列化到结构体
func (c *Client) GetStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.Get(ctx, key)
	})
}

// Set 设置键值（不过期）
func (c *Client) Set(ctx context.Context, key string, value interface{}) error {
	return c.UniversalClient.Set(ctx, key, value, 0).Err()
}

// SetEX 设置键值（带过期时间）
func (c *Client) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.UniversalClient.Set(ctx, key, value, expiration).Err()
}

// SetNX 仅在key不存在时设置（不过期）
func (c *Client) SetNX(ctx context.Context, key string, value interface{}) (bool, error) {
	return c.UniversalClient.SetNX(ctx, key, value, 0).Result()
}

// SetNXEX 仅在key不存在时设置（带过期时间）
func (c *Client) SetNXEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.UniversalClient.SetNX(ctx, key, value, expiration).Result()
}

// SetXX 仅在key存在时设置（不过期）
func (c *Client) SetXX(ctx context.Context, key string, value interface{}) (bool, error) {
	return c.UniversalClient.SetXX(ctx, key, value, 0).Result()
}

// SetXXEX 仅在key存在时设置（带过期时间）
func (c *Client) SetXXEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.UniversalClient.SetXX(ctx, key, value, expiration).Result()
}

// Append 追加字符串到键值
func (c *Client) Append(ctx context.Context, key, value string) (int64, error) {
	return c.UniversalClient.Append(ctx, key, value).Result()
}

// StrLen 获取字符串长度
func (c *Client) StrLen(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.StrLen(ctx, key).Result()
}

// GetRange 获取字符串的子串
func (c *Client) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	return c.UniversalClient.GetRange(ctx, key, start, end).Result()
}

// SetRange 设置字符串的子串
func (c *Client) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return c.UniversalClient.SetRange(ctx, key, offset, value).Result()
}

// Incr 递增键值
func (c *Client) Incr(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.Incr(ctx, key).Result()
}

// IncrBy 按指定值递增键值
func (c *Client) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.UniversalClient.IncrBy(ctx, key, value).Result()
}

// IncrByFloat 按浮点数值递增键值
func (c *Client) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return c.UniversalClient.IncrByFloat(ctx, key, value).Result()
}

// Decr 递减键值
func (c *Client) Decr(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.Decr(ctx, key).Result()
}

// DecrBy 按指定值递减键值
func (c *Client) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.UniversalClient.DecrBy(ctx, key, value).Result()
}

// GetSet 获取旧值并设置新值
func (c *Client) GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	return c.UniversalClient.GetSet(ctx, key, value).Result()
}

// GetSetStruct 获取旧值并设置新值，反序列化到结构体
func (c *Client) GetSetStruct(ctx context.Context, key string, value interface{}, result interface{}) error {
	return c.unmarshalValue(result, func() (string, error) {
		return c.GetSet(ctx, key, value)
	})
}

// MGet 批量获取键值
func (c *Client) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return c.UniversalClient.MGet(ctx, keys...).Result()
}

// MSet 批量设置键值
func (c *Client) MSet(ctx context.Context, values ...interface{}) error {
	return c.UniversalClient.MSet(ctx, values...).Err()
}

// MSetNX 批量设置键值（仅在所有key都不存在时设置）
func (c *Client) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	return c.UniversalClient.MSetNX(ctx, values...).Result()
}
