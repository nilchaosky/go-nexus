package viper

import (
	"fmt"

	"github.com/spf13/viper"
)

// Register 注册配置
func Register[T any](path string) (*T, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config T
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &config, nil
}
