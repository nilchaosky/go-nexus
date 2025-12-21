package log

import (
	"fmt"
	"os"

	"github.com/nilchaosky/go-nexus/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Logger 全局日志输出变量
	Logger *zap.Logger
)

// init 包加载时自动初始化日志
func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("日志初始化失败: " + err.Error())
	}
	Logger = logger
}

// Register 注册日志
func Register(config Config) error {
	// 确定输出目录，如果为空则使用默认目录
	outputDir := config.OutputDir
	if outputDir == "" {
		outputDir = "logs"
	}

	// 判断路径是否为目录
	if !utils.IsDir(outputDir) {
		// 路径不是目录，尝试创建
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("创建日志目录失败: %w", err)
		}
	}

	// 清理旧日志文件
	if err := cleanOldLogs(outputDir, config.MaxAge); err != nil {
		return fmt.Errorf("清理旧日志失败: %w", err)
	}

	// 获取编码器
	encoder := config.encoder()

	// 获取日志级别列表
	levels := config.getLevels()

	// 创建写入器（所有级别共享同一个文件）
	writeSyncer, err := createWriteSyncer(outputDir, config.ConsoleOutput)
	if err != nil {
		return err
	}

	// 创建多个 Core，每个级别一个
	cores := make([]zapcore.Core, 0, len(levels))
	for _, level := range levels {
		// 创建自定义 Core，处理该级别及以上的日志
		core := NewCustomCore(encoder, writeSyncer, level)
		cores = append(cores, core)
	}

	// 合并所有 Core
	core := zapcore.NewTee(cores...)

	// 创建日志选项
	var opts []zap.Option
	if config.ShowCaller {
		opts = append(opts, zap.AddCaller())
	}

	// 创建日志实例
	logger := zap.New(core, opts...)
	Logger = logger

	return nil
}
