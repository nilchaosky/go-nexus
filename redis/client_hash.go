package redis

import (
	"context"
	"reflect"

	"github.com/nilchaosky/go-nexus/nexus_utils"
	"github.com/nilchaosky/go-nexus/serialize"
)

// Hash 哈希表操作接口
type Hash interface {
	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HExists(ctx context.Context, key, field string) (bool, error)
	HGet(ctx context.Context, key, field string) (string, error)
	HGetStruct(ctx context.Context, key, field string, value interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)
	HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HLen(ctx context.Context, key string) (int64, error)
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	HMGetStruct(ctx context.Context, key string, value interface{}, fields ...string) error
	HMSet(ctx context.Context, key string, values ...interface{}) error
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error)
	HStrLen(ctx context.Context, key, field string) (int64, error)
	HVals(ctx context.Context, key string) ([]string, error)
	HValsStruct(ctx context.Context, key string, value interface{}) error
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}

// HDel 删除哈希表中的字段
func (c *Client) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return c.UniversalClient.HDel(ctx, key, fields...).Result()
}

// HExists 判断哈希表中字段是否存在
func (c *Client) HExists(ctx context.Context, key, field string) (bool, error) {
	return c.UniversalClient.HExists(ctx, key, field).Result()
}

// HGet 获取哈希表中字段的值
func (c *Client) HGet(ctx context.Context, key, field string) (string, error) {
	return c.UniversalClient.HGet(ctx, key, field).Result()
}

// HGetStruct 获取哈希表中字段的值，反序列化到结构体
func (c *Client) HGetStruct(ctx context.Context, key, field string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.HGet(ctx, key, field)
	})
}

// HGetAll 获取哈希表中所有字段和值
func (c *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.UniversalClient.HGetAll(ctx, key).Result()
}

// HIncrBy 将哈希表中字段的值增加指定整数
func (c *Client) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return c.UniversalClient.HIncrBy(ctx, key, field, incr).Result()
}

// HIncrByFloat 将哈希表中字段的值增加指定浮点数
func (c *Client) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	return c.UniversalClient.HIncrByFloat(ctx, key, field, incr).Result()
}

// HKeys 获取哈希表中所有字段
func (c *Client) HKeys(ctx context.Context, key string) ([]string, error) {
	return c.UniversalClient.HKeys(ctx, key).Result()
}

// HLen 获取哈希表中字段数量
func (c *Client) HLen(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.HLen(ctx, key).Result()
}

// HMGet 批量获取哈希表中字段的值
func (c *Client) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return c.UniversalClient.HMGet(ctx, key, fields...).Result()
}

// HMGetStruct 批量获取哈希表中字段的值，反序列化到结构体切片
func (c *Client) HMGetStruct(ctx context.Context, key string, value interface{}, fields ...string) error {
	_, err := nexus_utils.IsPointer(value)
	if err != nil {
		return err
	}

	results, err := c.HMGet(ctx, key, fields...)
	if err != nil {
		return err
	}

	rv, err := nexus_utils.IsSlice(value)
	if err != nil {
		return err
	}

	sliceType := rv.Type()
	elemType := sliceType.Elem()
	newSlice := reflect.MakeSlice(sliceType, 0, len(results))

	for _, result := range results {
		if result == nil {
			continue
		}

		str, ok := result.(string)
		if !ok || str == "" {
			continue
		}

		elemValue := reflect.New(elemType).Interface()
		if err := serialize.JSONIter.Unmarshal([]byte(str), elemValue); err != nil {
			return err
		}

		elemPtr := reflect.ValueOf(elemValue)
		newSlice = reflect.Append(newSlice, elemPtr.Elem())
	}

	rv.Set(newSlice)
	return nil
}

// HMSet 批量设置哈希表中字段的值
func (c *Client) HMSet(ctx context.Context, key string, values ...interface{}) error {
	return c.UniversalClient.HMSet(ctx, key, values...).Err()
}

// HSet 设置哈希表中字段的值
func (c *Client) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.UniversalClient.HSet(ctx, key, values...).Result()
}

// HSetNX 仅在字段不存在时设置哈希表中字段的值
func (c *Client) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return c.UniversalClient.HSetNX(ctx, key, field, value).Result()
}

// HStrLen 获取哈希表中字段值的字符串长度
func (c *Client) HStrLen(ctx context.Context, key, field string) (int64, error) {
	return c.UniversalClient.HStrLen(ctx, key, field).Result()
}

// HVals 获取哈希表中所有值
func (c *Client) HVals(ctx context.Context, key string) ([]string, error) {
	return c.UniversalClient.HVals(ctx, key).Result()
}

// HValsStruct 获取哈希表中所有值，反序列化到结构体切片
func (c *Client) HValsStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.HVals(ctx, key)
	})
}

// HScan 扫描哈希表
func (c *Client) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.UniversalClient.HScan(ctx, key, cursor, match, count).Result()
}
