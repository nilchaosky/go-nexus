package logz

import (
	"testing"
)

// TestRegister_FileNoColor 测试文件输出无颜色
func TestRegister_FileNoColor(t *testing.T) {
	config := Config{
		Level:         "info",
		Encoder:       "console",
		EncodeLevel:   "lowercase_color",
		Prefix:        "test",
		ConsoleOutput: true,
		ShowCaller:    true,
	}

	err := Register(config)
	if err != nil {
		t.Fatalf("注册日志失败: %v", err)
	}

	// 验证 Logger 已初始化
	if Logger == nil {
		t.Error("Logger 未初始化")
	}

	// 测试日志写入
	Logger.Info("文件输出测试（无颜色）")
	Logger.Warn("文件输出警告（无颜色）")
	Logger.Error("文件输出错误（无颜色）")
}
