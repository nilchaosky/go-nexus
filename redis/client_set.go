package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Set 集合操作接口
type Set interface {
	SAdd(ctx context.Context, key string, members ...interface{}) (int64, error)
	SCard(ctx context.Context, key string) (int64, error)
	SDiff(ctx context.Context, keys ...string) ([]string, error)
	SDiffStruct(ctx context.Context, value interface{}, keys ...string) error
	SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SInter(ctx context.Context, keys ...string) ([]string, error)
	SInterStruct(ctx context.Context, value interface{}, keys ...string) error
	SInterStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
	SMembers(ctx context.Context, key string) ([]string, error)
	SMembersStruct(ctx context.Context, key string, value interface{}) error
	SMove(ctx context.Context, source, destination string, member interface{}) (bool, error)
	SPop(ctx context.Context, key string) (string, error)
	SPopStruct(ctx context.Context, key string, value interface{}) error
	SPopN(ctx context.Context, key string, count int64) ([]string, error)
	SPopNStruct(ctx context.Context, key string, count int64, value interface{}) error
	SRandMember(ctx context.Context, key string) (string, error)
	SRandMemberStruct(ctx context.Context, key string, value interface{}) error
	SRandMemberN(ctx context.Context, key string, count int64) ([]string, error)
	SRandMemberNStruct(ctx context.Context, key string, count int64, value interface{}) error
	SRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	SUnion(ctx context.Context, keys ...string) ([]string, error)
	SUnionStruct(ctx context.Context, value interface{}, keys ...string) error
	SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error)
	SSCan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
}

// SAdd 向集合添加成员
func (c *Client) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.UniversalClient.SAdd(ctx, key, members...).Result()
}

// SCard 获取集合成员数量
func (c *Client) SCard(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.SCard(ctx, key).Result()
}

// SDiff 获取多个集合的差集
func (c *Client) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	return c.UniversalClient.SDiff(ctx, keys...).Result()
}

// SDiffStruct 获取多个集合的差集，反序列化到结构体切片
func (c *Client) SDiffStruct(ctx context.Context, value interface{}, keys ...string) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.SDiff(ctx, keys...)
	})
}

// SDiffStore 将多个集合的差集存储到目标集合
func (c *Client) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return c.UniversalClient.SDiffStore(ctx, destination, keys...).Result()
}

// SInter 获取多个集合的交集
func (c *Client) SInter(ctx context.Context, keys ...string) ([]string, error) {
	return c.UniversalClient.SInter(ctx, keys...).Result()
}

// SInterStruct 获取多个集合的交集，反序列化到结构体切片
func (c *Client) SInterStruct(ctx context.Context, value interface{}, keys ...string) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.SInter(ctx, keys...)
	})
}

// SInterStore 将多个集合的交集存储到目标集合
func (c *Client) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return c.UniversalClient.SInterStore(ctx, destination, keys...).Result()
}

// SIsMember 判断成员是否在集合中
func (c *Client) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return c.UniversalClient.SIsMember(ctx, key, member).Result()
}

// SMembers 获取集合所有成员
func (c *Client) SMembers(ctx context.Context, key string) ([]string, error) {
	return c.UniversalClient.SMembers(ctx, key).Result()
}

// SMembersStruct 获取集合所有成员，反序列化到结构体切片
func (c *Client) SMembersStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.SMembers(ctx, key)
	})
}

// SMove 将成员从一个集合移动到另一个集合
func (c *Client) SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	return c.UniversalClient.SMove(ctx, source, destination, member).Result()
}

// SPop 随机移除并返回集合中的一个成员
func (c *Client) SPop(ctx context.Context, key string) (string, error) {
	return c.UniversalClient.SPop(ctx, key).Result()
}

// SPopStruct 随机移除并返回集合中的一个成员，反序列化到结构体
func (c *Client) SPopStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.SPop(ctx, key)
	})
}

// SPopN 随机移除并返回集合中的N个成员
func (c *Client) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	return c.UniversalClient.SPopN(ctx, key, count).Result()
}

// SPopNStruct 随机移除并返回集合中的N个成员，反序列化到结构体切片
func (c *Client) SPopNStruct(ctx context.Context, key string, count int64, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.SPopN(ctx, key, count)
	})
}

// SRandMember 随机返回集合中的一个成员（不移除）
func (c *Client) SRandMember(ctx context.Context, key string) (string, error) {
	return c.UniversalClient.SRandMember(ctx, key).Result()
}

// SRandMemberStruct 随机返回集合中的一个成员（不移除），反序列化到结构体
func (c *Client) SRandMemberStruct(ctx context.Context, key string, value interface{}) error {
	return c.unmarshalValue(value, func() (string, error) {
		return c.SRandMember(ctx, key)
	})
}

// SRandMemberN 随机返回集合中的N个成员（不移除）
func (c *Client) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	return c.UniversalClient.SRandMemberN(ctx, key, count).Result()
}

// SRandMemberNStruct 随机返回集合中的N个成员（不移除），反序列化到结构体切片
func (c *Client) SRandMemberNStruct(ctx context.Context, key string, count int64, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.SRandMemberN(ctx, key, count)
	})
}

// SRem 从集合中移除成员
func (c *Client) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.UniversalClient.SRem(ctx, key, members...).Result()
}

// SUnion 获取多个集合的并集
func (c *Client) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	return c.UniversalClient.SUnion(ctx, keys...).Result()
}

// SUnionStruct 获取多个集合的并集，反序列化到结构体切片
func (c *Client) SUnionStruct(ctx context.Context, value interface{}, keys ...string) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.SUnion(ctx, keys...)
	})
}

// SUnionStore 将多个集合的并集存储到目标集合
func (c *Client) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return c.UniversalClient.SUnionStore(ctx, destination, keys...).Result()
}

// SSCan 扫描集合成员
func (c *Client) SSCan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return c.UniversalClient.SScan(ctx, key, cursor, match, count)
}
