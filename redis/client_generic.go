package redis

import (
	"context"
	"time"
)

// Generic 通用操作接口
type Generic interface {
	Del(ctx context.Context, keys ...string) (int64, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error)
	TTL(ctx context.Context, key string) (time.Duration, error)
	PTTL(ctx context.Context, key string) (time.Duration, error)
	Persist(ctx context.Context, key string) (bool, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)
	Type(ctx context.Context, key string) (string, error)
	Rename(ctx context.Context, key, newKey string) error
	RenameNX(ctx context.Context, key, newKey string) (bool, error)
	Move(ctx context.Context, key string, db int) (bool, error)
	RandomKey(ctx context.Context) (string, error)
	Dump(ctx context.Context, key string) (string, error)
	Restore(ctx context.Context, key string, ttl time.Duration, value string) error
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) error
}

// Del 删除一个或多个键
func (c *Client) Del(ctx context.Context, keys ...string) (int64, error) {
	return c.UniversalClient.Del(ctx, keys...).Result()
}

// Exists 检查一个或多个键是否存在
func (c *Client) Exists(ctx context.Context, keys ...string) (int64, error) {
	return c.UniversalClient.Exists(ctx, keys...).Result()
}

// Expire 设置键的过期时间（秒）
func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.UniversalClient.Expire(ctx, key, expiration).Result()
}

// ExpireAt 设置键的过期时间（Unix时间戳）
func (c *Client) ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return c.UniversalClient.ExpireAt(ctx, key, tm).Result()
}

// TTL 获取键的剩余过期时间（秒）
func (c *Client) TTL(ctx context.Context, key string) (time.Duration, error) {
	return c.UniversalClient.TTL(ctx, key).Result()
}

// PTTL 获取键的剩余过期时间（毫秒）
func (c *Client) PTTL(ctx context.Context, key string) (time.Duration, error) {
	return c.UniversalClient.PTTL(ctx, key).Result()
}

// Persist 移除键的过期时间，使其永久存在
func (c *Client) Persist(ctx context.Context, key string) (bool, error) {
	return c.UniversalClient.Persist(ctx, key).Result()
}

// Keys 获取所有匹配模式的键
func (c *Client) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.UniversalClient.Keys(ctx, pattern).Result()
}

// Scan 扫描键（推荐使用，避免阻塞）
func (c *Client) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.UniversalClient.Scan(ctx, cursor, match, count).Result()
}

// Type 获取键的类型
func (c *Client) Type(ctx context.Context, key string) (string, error) {
	return c.UniversalClient.Type(ctx, key).Result()
}

// Rename 重命名键
func (c *Client) Rename(ctx context.Context, key, newKey string) error {
	return c.UniversalClient.Rename(ctx, key, newKey).Err()
}

// RenameNX 仅在新键不存在时重命名键
func (c *Client) RenameNX(ctx context.Context, key, newKey string) (bool, error) {
	return c.UniversalClient.RenameNX(ctx, key, newKey).Result()
}

// Move 将键移动到指定数据库
func (c *Client) Move(ctx context.Context, key string, db int) (bool, error) {
	return c.UniversalClient.Move(ctx, key, db).Result()
}

// RandomKey 随机返回一个键
func (c *Client) RandomKey(ctx context.Context) (string, error) {
	return c.UniversalClient.RandomKey(ctx).Result()
}

// Dump 序列化键的值
func (c *Client) Dump(ctx context.Context, key string) (string, error) {
	return c.UniversalClient.Dump(ctx, key).Result()
}

// Restore 反序列化并恢复键的值
func (c *Client) Restore(ctx context.Context, key string, ttl time.Duration, value string) error {
	return c.UniversalClient.Restore(ctx, key, ttl, value).Err()
}

// RestoreReplace 反序列化并恢复键的值（如果键已存在则替换）
func (c *Client) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) error {
	return c.UniversalClient.RestoreReplace(ctx, key, ttl, value).Err()
}
