package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ZSet 有序集合操作接口
type ZSet interface {
	ZAdd(ctx context.Context, key string, members ...redis.Z) (int64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZCount(ctx context.Context, key, min, max string) (int64, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error)
	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) (int64, error)
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)
	ZPopMax(ctx context.Context, key string, count ...int64) ([]redis.Z, error)
	ZPopMin(ctx context.Context, key string, count ...int64) ([]redis.Z, error)
	ZRange(ctx context.Context, key string, start, end int64) ([]string, error)
	ZRangeStruct(ctx context.Context, key string, start, end int64, value interface{}) error
	ZRangeWithScores(ctx context.Context, key string, start, end int64) ([]redis.Z, error)
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error)
	ZRangeByScoreStruct(ctx context.Context, key string, opt *redis.ZRangeBy, value interface{}) error
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error)
	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error)
	ZRangeByLexStruct(ctx context.Context, key string, opt *redis.ZRangeBy, value interface{}) error
	ZRank(ctx context.Context, key, member string) (int64, error)
	ZRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	ZRemRangeByRank(ctx context.Context, key string, start, end int64) (int64, error)
	ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error)
	ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error)
	ZRevRange(ctx context.Context, key string, start, end int64) ([]string, error)
	ZRevRangeStruct(ctx context.Context, key string, start, end int64, value interface{}) error
	ZRevRangeWithScores(ctx context.Context, key string, start, end int64) ([]redis.Z, error)
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error)
	ZRevRangeByScoreStruct(ctx context.Context, key string, opt *redis.ZRangeBy, value interface{}) error
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error)
	ZRevRank(ctx context.Context, key, member string) (int64, error)
	ZScore(ctx context.Context, key, member string) (float64, error)
	ZUnionStore(ctx context.Context, destination string, store *redis.ZStore) (int64, error)
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}

// ZAdd 向有序集合添加成员
func (c *Client) ZAdd(ctx context.Context, key string, members ...redis.Z) (int64, error) {
	return c.UniversalClient.ZAdd(ctx, key, members...).Result()
}

// ZCard 获取有序集合成员数量
func (c *Client) ZCard(ctx context.Context, key string) (int64, error) {
	return c.UniversalClient.ZCard(ctx, key).Result()
}

// ZCount 统计有序集合中指定分数范围内的成员数量
func (c *Client) ZCount(ctx context.Context, key, min, max string) (int64, error) {
	return c.UniversalClient.ZCount(ctx, key, min, max).Result()
}

// ZIncrBy 将有序集合中成员的分数增加指定值
func (c *Client) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return c.UniversalClient.ZIncrBy(ctx, key, increment, member).Result()
}

// ZInterStore 将多个有序集合的交集存储到目标有序集合
func (c *Client) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) (int64, error) {
	return c.UniversalClient.ZInterStore(ctx, destination, store).Result()
}

// ZLexCount 统计有序集合中指定字典序范围内的成员数量
func (c *Client) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	return c.UniversalClient.ZLexCount(ctx, key, min, max).Result()
}

// ZPopMax 移除并返回有序集合中分数最高的成员
func (c *Client) ZPopMax(ctx context.Context, key string, count ...int64) ([]redis.Z, error) {
	return c.UniversalClient.ZPopMax(ctx, key, count...).Result()
}

// ZPopMin 移除并返回有序集合中分数最低的成员
func (c *Client) ZPopMin(ctx context.Context, key string, count ...int64) ([]redis.Z, error) {
	return c.UniversalClient.ZPopMin(ctx, key, count...).Result()
}

// ZRange 获取有序集合中指定范围的成员（按分数排序）
func (c *Client) ZRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	return c.UniversalClient.ZRange(ctx, key, start, end).Result()
}

// ZRangeStruct 获取有序集合中指定范围的成员（按分数排序），反序列化到结构体切片
func (c *Client) ZRangeStruct(ctx context.Context, key string, start, end int64, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.ZRange(ctx, key, start, end)
	})
}

// ZRangeWithScores 获取有序集合中指定范围的成员及其分数
func (c *Client) ZRangeWithScores(ctx context.Context, key string, start, end int64) ([]redis.Z, error) {
	return c.UniversalClient.ZRangeWithScores(ctx, key, start, end).Result()
}

// ZRangeByScore 获取有序集合中指定分数范围内的成员
func (c *Client) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.UniversalClient.ZRangeByScore(ctx, key, opt).Result()
}

// ZRangeByScoreStruct 获取有序集合中指定分数范围内的成员，反序列化到结构体切片
func (c *Client) ZRangeByScoreStruct(ctx context.Context, key string, opt *redis.ZRangeBy, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.ZRangeByScore(ctx, key, opt)
	})
}

// ZRangeByScoreWithScores 获取有序集合中指定分数范围内的成员及其分数
func (c *Client) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return c.UniversalClient.ZRangeByScoreWithScores(ctx, key, opt).Result()
}

// ZRangeByLex 获取有序集合中指定字典序范围内的成员
func (c *Client) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.UniversalClient.ZRangeByLex(ctx, key, opt).Result()
}

// ZRangeByLexStruct 获取有序集合中指定字典序范围内的成员，反序列化到结构体切片
func (c *Client) ZRangeByLexStruct(ctx context.Context, key string, opt *redis.ZRangeBy, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.ZRangeByLex(ctx, key, opt)
	})
}

// ZRank 获取有序集合中成员的排名（从0开始，分数从低到高）
func (c *Client) ZRank(ctx context.Context, key, member string) (int64, error) {
	return c.UniversalClient.ZRank(ctx, key, member).Result()
}

// ZRem 从有序集合中移除成员
func (c *Client) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return c.UniversalClient.ZRem(ctx, key, members...).Result()
}

// ZRemRangeByRank 移除有序集合中指定排名范围内的成员
func (c *Client) ZRemRangeByRank(ctx context.Context, key string, start, end int64) (int64, error) {
	return c.UniversalClient.ZRemRangeByRank(ctx, key, start, end).Result()
}

// ZRemRangeByScore 移除有序集合中指定分数范围内的成员
func (c *Client) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	return c.UniversalClient.ZRemRangeByScore(ctx, key, min, max).Result()
}

// ZRemRangeByLex 移除有序集合中指定字典序范围内的成员
func (c *Client) ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error) {
	return c.UniversalClient.ZRemRangeByLex(ctx, key, min, max).Result()
}

// ZRevRange 获取有序集合中指定范围的成员（按分数从高到低排序）
func (c *Client) ZRevRange(ctx context.Context, key string, start, end int64) ([]string, error) {
	return c.UniversalClient.ZRevRange(ctx, key, start, end).Result()
}

// ZRevRangeStruct 获取有序集合中指定范围的成员（按分数从高到低排序），反序列化到结构体切片
func (c *Client) ZRevRangeStruct(ctx context.Context, key string, start, end int64, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.ZRevRange(ctx, key, start, end)
	})
}

// ZRevRangeWithScores 获取有序集合中指定范围的成员及其分数（按分数从高到低排序）
func (c *Client) ZRevRangeWithScores(ctx context.Context, key string, start, end int64) ([]redis.Z, error) {
	return c.UniversalClient.ZRevRangeWithScores(ctx, key, start, end).Result()
}

// ZRevRangeByScore 获取有序集合中指定分数范围内的成员（按分数从高到低排序）
func (c *Client) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.UniversalClient.ZRevRangeByScore(ctx, key, opt).Result()
}

// ZRevRangeByScoreStruct 获取有序集合中指定分数范围内的成员（按分数从高到低排序），反序列化到结构体切片
func (c *Client) ZRevRangeByScoreStruct(ctx context.Context, key string, opt *redis.ZRangeBy, value interface{}) error {
	return c.unmarshalSlice(value, func() ([]string, error) {
		return c.ZRevRangeByScore(ctx, key, opt)
	})
}

// ZRevRangeByScoreWithScores 获取有序集合中指定分数范围内的成员及其分数（按分数从高到低排序）
func (c *Client) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return c.UniversalClient.ZRevRangeByScoreWithScores(ctx, key, opt).Result()
}

// ZRevRank 获取有序集合中成员的排名（从0开始，分数从高到低）
func (c *Client) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	return c.UniversalClient.ZRevRank(ctx, key, member).Result()
}

// ZScore 获取有序集合中成员的分数
func (c *Client) ZScore(ctx context.Context, key, member string) (float64, error) {
	return c.UniversalClient.ZScore(ctx, key, member).Result()
}

// ZUnionStore 将多个有序集合的并集存储到目标有序集合
func (c *Client) ZUnionStore(ctx context.Context, destination string, store *redis.ZStore) (int64, error) {
	return c.UniversalClient.ZUnionStore(ctx, destination, store).Result()
}

// ZScan 扫描有序集合
func (c *Client) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.UniversalClient.ZScan(ctx, key, cursor, match, count).Result()
}
