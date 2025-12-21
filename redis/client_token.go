package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nilchaosky/go-nexus/redis/token"
)

var (
	userTokenRedisKeyPrefix = "USER:"
	tokenRedisKey           = "Token"
	refreshTokenRedisKey    = "RefreshKey"
)

// Token Token操作接口
type Token interface {
	GetUserTokenKey(id string) (string, error)
	GetToken(ctx context.Context, id string) (string, string)
	SaveToken(ctx context.Context, id string, config token.Config, extra map[string]interface{}) error
	DeleteToken(ctx context.Context, id string) error
	RefreshToken(ctx context.Context, id string, oldToken string, oldRefreshToken string, config token.Config) error
}

// GetUserTokenKey 获取用户Token Key
func (c *Client) GetUserTokenKey(id string) (string, error) {
	if id == "" {
		return "", errors.New("id不能为空")
	}
	return userTokenRedisKeyPrefix + id, nil
}

// GetToken 获取Token和RefreshToken
func (c *Client) GetToken(ctx context.Context, id string) (string, string) {
	// 获取用户Token Key
	key, err := c.GetUserTokenKey(id)
	if err != nil {
		return "", ""
	}
	tokenKey := key + ":" + tokenRedisKey
	refreshTokenKey := key + ":" + refreshTokenRedisKey

	// 从Redis获取Token
	tokenValue, err := c.Get(ctx, tokenKey)
	if err != nil {
		return "", ""
	}

	// 从Redis获取RefreshToken
	refreshTokenValue, err := c.Get(ctx, refreshTokenKey)
	if err != nil {
		return "", ""
	}

	return tokenValue, refreshTokenValue
}

// SaveToken 保存Token到Redis
func (c *Client) SaveToken(ctx context.Context, id string, config token.Config, extra map[string]interface{}) error {
	// 获取用户Token Key
	key, err := c.GetUserTokenKey(id)
	if err != nil {
		return err
	}
	tokenKey := key + ":" + tokenRedisKey
	refreshTokenKey := key + ":" + refreshTokenRedisKey

	// 将id添加到extra中
	if extra == nil {
		extra = make(map[string]interface{})
	}
	extra["id"] = id

	// 使用Generate生成Token
	tokenValue, refreshTokenValue, err := token.Generate(config, extra)
	if err != nil {
		return fmt.Errorf("生成Token失败: %w", err)
	}

	// 使用config中的过期时间（将天数转换为Duration）
	expiration := time.Duration(config.Expiration) * 24 * time.Hour

	// 保存Token和RefreshToken到Redis
	if err := c.SetEX(ctx, tokenKey, tokenValue, expiration).Err(); err != nil {
		return fmt.Errorf("保存Token失败: %w", err)
	}
	if err := c.SetEX(ctx, refreshTokenKey, refreshTokenValue, expiration).Err(); err != nil {
		return fmt.Errorf("保存RefreshToken失败: %w", err)
	}

	return nil
}

// DeleteToken 删除Token
func (c *Client) DeleteToken(ctx context.Context, id string) error {
	// 获取用户Token Key
	key, err := c.GetUserTokenKey(id)
	if err != nil {
		return err
	}

	// 构建Token和RefreshToken的key
	tokenKey := key + ":" + tokenRedisKey
	refreshTokenKey := key + ":" + refreshTokenRedisKey

	// 删除Token和RefreshToken
	_, err = c.Del(ctx, tokenKey, refreshTokenKey)
	if err != nil {
		return fmt.Errorf("删除Token失败: %w", err)
	}

	return nil
}

// RefreshToken 刷新Token
func (c *Client) RefreshToken(ctx context.Context, id string, oldToken string, oldRefreshToken string, config token.Config) error {
	// 验证oldToken和oldRefreshToken是否为空
	if oldToken == "" || oldRefreshToken == "" {
		return errors.New("token丢失")
	}

	// 从旧的Token中获取额外参数
	extra, err := token.GetExtra(config, oldToken)
	if err != nil {
		return errors.New("token解析失败")
	}

	// 对比extra中的id和参数id是否一致
	if extraID, ok := extra["id"].(string); !ok || extraID != id {
		return errors.New("无效的Token")
	}

	// 获取用户Token Key
	key, err := c.GetUserTokenKey(id)
	if err != nil {
		return err
	}
	tokenKey := key + ":" + tokenRedisKey
	refreshTokenKey := key + ":" + refreshTokenRedisKey

	// 验证旧的oldToken
	if err := token.Verify(config, oldToken); err != nil {
		return fmt.Errorf("刷新Token验证失败: %w", err)
	}

	// 从Redis获取Token和RefreshToken
	_, refreshTokenValue := c.GetToken(ctx, id)

	// 验证刷新Token是否一致
	if refreshTokenValue != oldRefreshToken {
		return errors.New("刷新Token无效")
	}

	// 生成新的Token和RefreshToken
	tokenValue, refreshTokenValue, err := token.Generate(config, extra)
	if err != nil {
		return fmt.Errorf("生成Token失败: %w", err)
	}

	// 使用config中的过期时间（将天数转换为Duration）
	expiration := time.Duration(config.Expiration) * 24 * time.Hour

	// 保存新的Token和RefreshToken到Redis
	if err := c.SetEX(ctx, tokenKey, tokenValue, expiration).Err(); err != nil {
		return fmt.Errorf("保存Token失败: %w", err)
	}
	if err := c.SetEX(ctx, refreshTokenKey, refreshTokenValue, expiration).Err(); err != nil {
		return fmt.Errorf("保存RefreshToken失败: %w", err)
	}

	return nil
}
