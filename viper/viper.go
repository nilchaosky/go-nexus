package viper

import (
	"fmt"

	"github.com/spf13/viper"
)

// Register 注册配置
func Register[T any](path string, config *T) error {
	if config == nil {
		return fmt.Errorf("config不能为nil")
	}

	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 直接反序列化到config
	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	return nil
}
