package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用bcrypt对密码进行哈希
// password 原始密码
// cost bcrypt成本因子，范围4-31，默认10
// 返回哈希后的密码字符串和错误
func HashPassword(password string, cost ...int) (string, error) {
	if password == "" {
		return "", errors.New("密码不能为空")
	}

	// 默认成本因子为10
	bcryptCost := 10
	if len(cost) > 0 && cost[0] > 0 {
		bcryptCost = cost[0]
		// 限制成本因子范围
		if bcryptCost < 4 {
			bcryptCost = 4
		}
		if bcryptCost > 31 {
			bcryptCost = 31
		}
	}

	// 生成哈希
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// ComparePassword 比较原始密码和哈希密码是否匹配
// password 原始密码
// hashedPassword 哈希后的密码
// 返回匹配返回nil，不匹配返回错误
func ComparePassword(password, hashedPassword string) error {
	if password == "" {
		return errors.New("密码不能为空")
	}
	if hashedPassword == "" {
		return errors.New("哈希密码不能为空")
	}

	// 比较密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("密码不匹配")
	}

	return nil
}
