package redis

import (
	"context"
	"errors"
	"time"

	"github.com/nilchaosky/go-nexus/nexusutils"
	"github.com/nilchaosky/go-nexus/serialize"
	"github.com/redis/go-redis/v9"
)

// List 列表操作接口
type List interface {
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	BLPopStruct(ctx context.Context, timeout time.Duration, value interface{}, keys ...string) error
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	BRPopStruct(ctx context.Context, timeout time.Duration, value interface{}, keys ...string) error
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error)
	BRPopLPushStruct(ctx context.Context, source, destination string, timeout time.Duration, value interface{}) error
	LIndex(ctx context.Context, key string, index int64) (string, error)
	LIndexStruct(ctx context.Context, key string, index int64, value interface{}) error
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error)
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	LPop(ctx context.Context, key string) (string, error)
	LPopStruct(ctx context.Context, key string, value interface{}) error
	LPopCount(ctx context.Context, key string, count int) ([]string, error)
	LPopCountStruct(ctx context.Context, key string, count int, value interface{}) error
	LPos(ctx context.Context, key string, value string, args redis.LPosArgs) (int64, error)
	LPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	LPushX(ctx context.Context, key string, values ...interface{}) (int64, error)
	LRange(ctx context.Context, key string, start, end int64) ([]string, error)
	LRangeStruct(ctx context.Context, key string, start, end int64, value interface{}) error
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) error
	LTrim(ctx context.Context, key string, start, end int64) error
	RPop(ctx context.Context, key string) (string, error)
	RPopStruct(ctx context.Context, key string, value interface{}) error
	RPopCount(ctx context.Context, key string, count int) ([]string, error)
	RPopCountStruct(ctx context.Context, key string, count int, value interface{}) error
	RPopLPush(ctx context.Context, source, destination string) (string, error)
	RPopLPushStruct(ctx context.Context, source, destination string, value interface{}) error
	RPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	RPushX(ctx context.Context, key string, values ...interface{}) (int64, error)
}

// BLPop 阻塞式从列表左侧弹出元素
func (c *Client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return c.UniversalClient.BLPop(ctx, timeout, keys...).Result()
}

// BLPopStruct 阻塞式从列表左侧弹出元素，反序列化到结构体
func (c *Client) BLPopStruct(ctx context.Context, timeout time.Duration, value interface{}, keys ...string) error {
	_, err := nexus_utils.IsPointer(value)
	if err != nil {
		return err
	}

	results, err := c.BLPop(ctx, timeout, keys...)
	if err != nil {
		return err
	}

	if len(results) < 2 {
		return errors.New("结果格式错误")
	}

	data := results[1]
	if data == "" {
		return errors.New("键值为空")
	}

	if err := serialize.JSONIter.Unmarshal([]byte(data), value); err != nil {
		return err
	}

	return nil
}

// BRPop 阻塞式从列表右侧弹出元素
func (c *Client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return c.UniversalClient.BRPop(ctx, timeout, keys...).Result()
}

// BRPopStruct 阻塞式从列表右侧弹出元素，反序列化到结构体
func (c *Client) BRPopStruct(ctx context.Context, timeout time.Duration, value interface{}, keys ...string) error {
	_, err := nexus_utils.IsPointer(value)
	if err != nil {
		return err
	}

	results, err := c.BRPop(ctx, timeout, keys...)
	if err != nil {
		return err
	}

	if len(results) < 2 {
		return errors.New("结果格式错误")
	}

	data := results[1]
	if data == "" {
		return errors.New("键值为空")
	}

	if err := serialize.JSONIter.Unmarshal([]byte(data), value); err != nil {
		return err
	}

	return nil
}

// BRPopLPush 阻塞式从源列表右侧弹出元素并推入目标列表左侧
func (c *Client) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	return c.UniversalClient.BRPopLPush(ctx, source, destination, timeout).Result()
}

// BRPopLPushStruct 阻塞式从源列表右侧弹出元素并推入目标列表左侧，反序列化到结构体
func (c *Client) BRPopLPushStruct(ctx context.Context, source, destination string, timeout time.Duration, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.BRPopLPush(ctx, source, destination, timeout)
	})
}

// LIndex 获取列表中指定索引的元素
func (c *Client) LIndex(ctx context.Context, key string, index int64) (string, error) {
	return c.UniversalClient.LIndex(ctx, key, index).Result()
}

// LIndexStruct 获取列表中指定索引的元素并反序列化到结构体
func (c *Client) LIndexStruct(ctx context.Context, key string, index int64, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.LIndex(ctx, key, index)
	})
}

// LInsert 在列表的指定元素前或后插入新元素
func (c *Client) LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	return c.UniversalClient.LInsert(ctx, key, op, pivot, value).Result()
}

// LInsertBefore 在列表的指定元素前插入新元素
func (c *Client) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return c.LInsert(ctx, key, "BEFORE", pivot, value)
}

// LInsertAfter 在列表的指定元素后插入新元素
func (c *Client) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return c.LInsert(ctx, key, "AFTER", pivot, value)
}

// LLen 获取列表长度
func (c *Client) LLen(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.LLen(ctx, key).Result()
}

// LPop 从列表左侧弹出元素
func (c *Client) LPop(ctx context.Context, key string) (string, error) {
	return c.UniversalClient.LPop(ctx, key).Result()
}

// LPopStruct 从列表左侧弹出元素并反序列化到结构体
func (c *Client) LPopStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.LPop(ctx, key)
	})
}

// LPopCount 从列表左侧弹出N个元素
func (c *Client) LPopCount(ctx context.Context, key string, count int) ([]string, error) {
	return c.UniversalClient.LPopCount(ctx, key, count).Result()
}

// LPopCountStruct 从列表左侧弹出N个元素并反序列化到结构体切片
func (c *Client) LPopCountStruct(ctx context.Context, key string, count int, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.LPopCount(ctx, key, count)
	})
}

// LPos 获取列表中元素的位置
func (c *Client) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) (int64, error) {
	return c.UniversalClient.LPos(ctx, key, value, args).Result()
}

// LPush 从列表左侧推入元素
func (c *Client) LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.UniversalClient.LPush(ctx, key, values...).Result()
}

// LPushX 仅在列表存在时从左侧推入元素
func (c *Client) LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.UniversalClient.LPushX(ctx, key, values...).Result()
}

// LRange 获取列表中指定范围的元素
func (c *Client) LRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	return c.UniversalClient.LRange(ctx, key, start, end).Result()
}

// LRangeStruct 获取列表中指定范围的元素并反序列化到结构体切片
func (c *Client) LRangeStruct(ctx context.Context, key string, start, end int64, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.LRange(ctx, key, start, end)
	})
}

// LRem 从列表中移除指定元素
func (c *Client) LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return c.UniversalClient.LRem(ctx, key, count, value).Result()
}

// LSet 设置列表中指定索引的元素
func (c *Client) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	return c.UniversalClient.LSet(ctx, key, index, value).Err()
}

// LTrim 修剪列表，只保留指定范围的元素
func (c *Client) LTrim(ctx context.Context, key string, start, end int64) error {
	return c.UniversalClient.LTrim(ctx, key, start, end).Err()
}

// RPop 从列表右侧弹出元素
func (c *Client) RPop(ctx context.Context, key string) (string, error) {
	return c.UniversalClient.RPop(ctx, key).Result()
}

// RPopStruct 从列表右侧弹出元素并反序列化到结构体
func (c *Client) RPopStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.RPop(ctx, key)
	})
}

// RPopCount 从列表右侧弹出N个元素
func (c *Client) RPopCount(ctx context.Context, key string, count int) ([]string, error) {
	return c.UniversalClient.RPopCount(ctx, key, count).Result()
}

// RPopCountStruct 从列表右侧弹出N个元素并反序列化到结构体切片
func (c *Client) RPopCountStruct(ctx context.Context, key string, count int, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.RPopCount(ctx, key, count)
	})
}

// RPopLPush 从源列表右侧弹出元素并推入目标列表左侧
func (c *Client) RPopLPush(ctx context.Context, source, destination string) (string, error) {
	return c.UniversalClient.RPopLPush(ctx, source, destination).Result()
}

// RPopLPushStruct 从源列表右侧弹出元素并推入目标列表左侧，反序列化到结构体
func (c *Client) RPopLPushStruct(ctx context.Context, source, destination string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.RPopLPush(ctx, source, destination)
	})
}

// RPush 从列表右侧推入元素
func (c *Client) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.UniversalClient.RPush(ctx, key, values...).Result()
}

// RPushX 仅在列表存在时从右侧推入元素
func (c *Client) RPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return c.UniversalClient.RPushX(ctx, key, values...).Result()
}
