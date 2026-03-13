package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	// Client 全局客户端
	Client *client
)

// Register 注册客户端
func Register(config Config) error {
	// 验证地址和端口
	if config.Address == "" {
		return errors.New("地址不能为空")
	}

	if config.Port == 0 {
		return errors.New("端口不能为空")
	}

	ctx := context.Background()

	// 如果DB未设置，默认使用0
	db := config.DB
	if db < 0 {
		return fmt.Errorf("数据库编号不能小于0，当前值: %d", db)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Address, config.Port),
		Password: config.Password,
		DB:       db,
	})

	// 测试连接，失败则返回错误
	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("数据库连接失败: %w", err)
	}

	// 将客户端存储到全局变量
	Client = NewClient(rdb)

	return nil
}
