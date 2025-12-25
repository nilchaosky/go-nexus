package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generate 生成Token
func Generate(config Config, id string, extra map[string]interface{}) (string, string, error) {
	if err := config.validate(); err != nil {
		return "", "", err
	}

	now := time.Now()
	expiresAt := now.AddDate(0, 0, config.Expiration)

	// 创建Access Token Claims
	accessClaims := jwt.MapClaims{
		"iss": config.Issuer,
		"iat": now.Unix(),
		"exp": expiresAt.Unix(),
		"id":  id,
	}

	// 将额外参数添加到Claims
	for k, v := range extra {
		accessClaims[k] = v
	}

	// 创建Access Token
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// 签名Access Token
	accessToken, err := access.SignedString([]byte(config.Secret))
	if err != nil {
		return "", "", fmt.Errorf("签名Token失败: %w", err)
	}

	// 格式化过期时间为字符串
	refreshToken := expiresAt.Format(time.DateTime)

	return accessToken, refreshToken, nil
}

// Verify 验证Token
func Verify(secret, tokenString string) error {
	if secret == "" {
		return errors.New("密钥不能为空")
	}

	// 解析Token，自动验证过期时间
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}, jwt.WithExpirationRequired())

	if err != nil {
		return fmt.Errorf("解析Token失败: %w", err)
	}

	// 验证Token是否有效（包括签名和过期时间）
	if !token.Valid {
		return errors.New("token无效")
	}

	return nil
}

// GetExtra 获取Token中的额外参数
func GetExtra(secret, tokenString string) (map[string]interface{}, error) {
	if secret == "" {
		return nil, errors.New("密钥不能为空")
	}

	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("不支持的签名方法: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("解析Token失败: %w", err)
	}

	// 验证Token是否有效
	if !token.Valid {
		return nil, errors.New("token无效")
	}

	// 获取Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token Claims格式错误")
	}

	// 移除标准Claims字段，只保留额外参数
	extraMap := make(map[string]interface{})
	standardFields := map[string]bool{
		"iss": true,
		"iat": true,
		"exp": true,
	}

	for k, v := range claims {
		if !standardFields[k] {
			extraMap[k] = v
		}
	}

	// 如果没有额外参数，返回空map
	if len(extraMap) == 0 {
		return make(map[string]interface{}), nil
	}

	return extraMap, nil
}
