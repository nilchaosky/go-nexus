package viper

import (
	"fmt"
	"reflect"

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

	// 读取配置文件到临时变量
	var fileConfig T
	if err := v.Unmarshal(&fileConfig); err != nil {
		return fmt.Errorf("解析配置失败: %w", err)
	}

	// 使用反射合并配置，只替换零值字段
	if err := mergeConfig(config, &fileConfig); err != nil {
		return fmt.Errorf("合并配置失败: %w", err)
	}

	return nil
}

// mergeConfig 合并配置，只替换零值字段
func mergeConfig(dst, src interface{}) error {
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)

	// 确保都是指针类型
	if dstValue.Kind() != reflect.Ptr || srcValue.Kind() != reflect.Ptr {
		return fmt.Errorf("参数必须是指针类型")
	}

	// 获取元素值
	dstElem := dstValue.Elem()
	srcElem := srcValue.Elem()

	// 确保都是结构体类型
	if dstElem.Kind() != reflect.Struct || srcElem.Kind() != reflect.Struct {
		return fmt.Errorf("参数必须是结构体类型")
	}

	// 遍历所有字段
	dstType := dstElem.Type()
	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstElem.Field(i)
		srcField := srcElem.Field(i)

		// 如果目标字段是零值，则用源字段的值替换
		if dstField.IsZero() && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}

	return nil
}
