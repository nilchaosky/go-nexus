package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap/zapcore"
)

// customCore 自定义 Core 实现
type customCore struct {
	encoder  zapcore.Encoder
	writer   zapcore.WriteSyncer
	minLevel zapcore.Level
}

// newCustomCore 创建自定义 Core
func newCustomCore(encoder zapcore.Encoder, writer zapcore.WriteSyncer, minLevel zapcore.Level) *customCore {
	return &customCore{
		encoder:  encoder,
		writer:   writer,
		minLevel: minLevel,
	}
}

// Enabled 判断级别是否启用
func (c *customCore) Enabled(level zapcore.Level) bool {
	return level >= c.minLevel
}

// With 添加字段
func (c *customCore) With(fields []zapcore.Field) zapcore.Core {
	encoder := c.encoder.Clone()
	for _, field := range fields {
		field.AddTo(encoder)
	}
	return &customCore{
		encoder:  encoder,
		writer:   c.writer,
		minLevel: c.minLevel,
	}
}

// Check 检查并准备写入
func (c *customCore) Check(entry zapcore.Entry, checkedEntry *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checkedEntry.AddCore(entry, c)
	}
	return checkedEntry
}

// Write 写入日志
func (c *customCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.encoder.EncodeEntry(entry, fields)
	if err != nil {
		return err
	}
	_, err = c.writer.Write(buf.Bytes())
	buf.Free()
	return err
}

// Sync 同步写入
func (c *customCore) Sync() error {
	return c.writer.Sync()
}

// createWriteSyncer 创建写入器
func createWriteSyncer(outputDir string, consoleOutput bool) (zapcore.WriteSyncer, error) {
	dateStr := time.Now().Format(time.DateOnly)
	filePath := filepath.Join(outputDir, dateStr+".log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件失败: %w", err)
	}
	writeSyncer := zapcore.AddSync(file)

	if consoleOutput {
		consoleSyncer := zapcore.AddSync(os.Stdout)
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, consoleSyncer)
	}

	return writeSyncer, nil
}

// cleanOldLogs 清理旧日志文件
func cleanOldLogs(outputDir string, maxAge int) error {
	if maxAge <= 0 {
		return nil
	}

	cutoffTime := time.Now().AddDate(0, 0, -maxAge)
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		return fmt.Errorf("读取日志目录失败: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !isLogFile(name) {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.ModTime().Before(cutoffTime) {
			filePath := filepath.Join(outputDir, name)
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("删除日志文件失败: %w", err)
			}
		}
	}

	return nil
}

// isLogFile 判断是否为日志文件
func isLogFile(name string) bool {
	if len(name) < 11 {
		return false
	}
	if name[len(name)-4:] != ".log" {
		return false
	}
	dateStr := name[:len(name)-4]
	_, err := time.Parse(time.DateOnly, dateStr)
	return err == nil
}
