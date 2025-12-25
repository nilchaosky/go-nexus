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
	SaveToken(ctx context.Context, id, tokenValue, refreshTokenValue string, expiration, refreshExpiration time.Duration) error
	DeleteToken(ctx context.Context, id string) error
	VerifyRefreshToken(ctx context.Context, id, oldToken, oldRefreshToken, secret string) error
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

	// 获取Token
	tokenValue, err := c.Get(ctx, tokenKey)
	if err != nil {
		return "", ""
	}

	// 获取RefreshToken
	refreshTokenValue, err := c.Get(ctx, refreshTokenKey)
	if err != nil {
		return "", ""
	}

	return tokenValue, refreshTokenValue
}

// SaveToken 保存Token
func (c *Client) SaveToken(ctx context.Context, id, tokenValue, refreshTokenValue string, expiration, refreshExpiration time.Duration) error {
	// 获取用户Token Key
	key, err := c.GetUserTokenKey(id)
	if err != nil {
		return err
	}
	tokenKey := key + ":" + tokenRedisKey
	refreshTokenKey := key + ":" + refreshTokenRedisKey

	// 保存Token和RefreshToken，使用不同的过期时间
	if err := c.SetEX(ctx, tokenKey, tokenValue, expiration); err != nil {
		return fmt.Errorf("保存Token失败: %w", err)
	}
	if err := c.SetEX(ctx, refreshTokenKey, refreshTokenValue, refreshExpiration); err != nil {
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

// VerifyRefreshToken 验证刷新Token
func (c *Client) VerifyRefreshToken(ctx context.Context, id, oldToken, oldRefreshToken, secret string) error {
	// 验证oldToken和oldRefreshToken是否为空
	if oldToken == "" || oldRefreshToken == "" {
		return errors.New("token丢失")
	}

	// 从旧的Token中获取额外参数
	extra, err := token.GetExtra(secret, oldToken)
	if err != nil {
		return errors.New("token解析失败")
	}

	// 对比extra中的id和参数id是否一致
	if extraID, ok := extra["id"].(string); !ok || extraID != id {
		return errors.New("无效的Token")
	}

	// 验证旧的oldToken
	if err := token.Verify(secret, oldToken); err != nil {
		return fmt.Errorf("刷新Token验证失败: %w", err)
	}

	// 获取Token和RefreshToken
	_, refreshTokenValue := c.GetToken(ctx, id)

	// 验证刷新Token是否一致
	if refreshTokenValue != oldRefreshToken {
		return errors.New("刷新Token无效")
	}

	return nil
}
