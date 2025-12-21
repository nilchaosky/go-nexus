package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/nilchaosky/go-nexus/log"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	// Clients Redis客户端列表
	Clients []*Client
	// index 默认客户端索引（默认：0）
	index int
)

// init 包加载时自动初始化
func init() {
	index = 0
	// Redis默认最大数据库数量为16（0-15）
	Clients = make([]*Client, 16)
}

// Register 注册Redis客户端
func Register(config Config) error {
	// 验证地址和端口
	if config.Address == "" {
		return errors.New("redis地址不能为空")
	}

	if config.Port == 0 {
		return errors.New("redis端口不能为空")
	}

	ctx := context.Background()

	// 如果DB列表为空，直接注册到索引0，连接失败则报错
	if len(config.DB) == 0 {
		client := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", config.Address, config.Port),
			DB:   0,
		})

		// 测试连接，失败则报错
		if err := client.Ping(ctx).Err(); err != nil {
			return fmt.Errorf("redis数据库0连接失败: %w", err)
		}

		Clients[0] = NewClient(client)
		return nil
	}

	// 如果DB列表不为空，循环处理，连接失败则跳过
	for _, db := range config.DB {
		// 验证数据库编号范围，超出范围则跳过
		if db < 0 || db >= len(Clients) {
			log.Logger.Warn("redis数据库超出范围，已跳过",
				zap.Int("db", db),
				zap.Int("min", 0),
				zap.Int("max", len(Clients)-1),
			)
			continue
		}

		client := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", config.Address, config.Port),
			DB:   db,
		})

		// 测试连接，失败则跳过
		if err := client.Ping(ctx).Err(); err != nil {
			log.Logger.Warn("redis数据库连接失败，已跳过",
				zap.Int("db", db),
				zap.Error(err),
			)
			continue
		}

		// 将客户端存储到对应索引位置
		Clients[db] = NewClient(client)
	}

	return nil
}

// GetClient 获取指定数据库的Redis客户端
func GetClient(db int) (*Client, error) {
	if db < 0 || db >= len(Clients) {
		return nil, fmt.Errorf("数据库编号%d超出范围（0-%d）", db, len(Clients)-1)
	}

	client := Clients[db]
	if client == nil {
		return nil, fmt.Errorf("数据库%d的客户端不存在", db)
	}

	return client, nil
}

// SetIndex 设置默认客户端索引
func SetIndex(i int) error {
	if i < 0 || i >= len(Clients) {
		return errors.New("redis数据库超出范围")
	}
	index = i
	return nil
}

// GetDefaultClient 获取默认Redis客户端（使用index索引）
func GetDefaultClient() (*Client, error) {
	client := Clients[index]
	if client == nil {
		return nil, fmt.Errorf("索引%d的客户端未初始化", index)
	}

	return client, nil
}
