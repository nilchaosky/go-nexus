package log

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// Config 日志配置结构体
type Config struct {
	Level         string `json:"level"`          // 日志等级：debug, info, warn, error（默认：debug）
	Encoder       string `json:"encoder"`        // 编码器类型：json, console（默认：console）
	EncodeLevel   string `json:"encode_level"`   // 日志级别编码器：lowercase, lowercase_color, capital, capital_color（默认：lowercase）
	Prefix        string `json:"prefix"`         // 日志前缀（默认：空）
	OutputDir     string `json:"output_dir"`     // 日志输出目录（默认：logs）
	ShowCaller    bool   `json:"show_caller"`    // 是否显示文件名和行号（默认：false）
	ConsoleOutput bool   `json:"console_output"` // 是否同时输出到控制台（默认：false）
	MaxAge        int    `json:"max_age"`        // 日志保留天数（默认：0，不清理）
}

// getLevels 返回从指定级别到最高级别的所有级别
func (c *Config) getLevels() []zapcore.Level {
	minLevel, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		minLevel = zapcore.DebugLevel
	}

	result := make([]zapcore.Level, 0, int(zapcore.FatalLevel-minLevel)+1)
	for l := minLevel; l <= zapcore.FatalLevel; l++ {
		result = append(result, l)
	}

	return result
}

// levelEncoder 根据配置返回日志级别编码器
func (c *Config) levelEncoder() zapcore.LevelEncoder {
	switch c.EncodeLevel {
	case "lowercase_color":
		return zapcore.LowercaseColorLevelEncoder
	case "capital":
		return zapcore.CapitalLevelEncoder
	case "capital_color":
		return zapcore.CapitalColorLevelEncoder
	case "lowercase":
		fallthrough
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// timeEncoder 返回时间编码器
func (c *Config) timeEncoder() zapcore.TimeEncoder {
	return func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		timeStr := t.Format(time.DateTime)
		if c.Prefix != "" {
			encoder.AppendString("[" + c.Prefix + "] " + timeStr)
		} else {
			encoder.AppendString(timeStr)
		}
	}
}

// encoder 创建编码器
func (c *Config) encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    c.levelEncoder(),
		EncodeTime:     c.timeEncoder(),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	switch c.Encoder {
	case "json":
		return zapcore.NewJSONEncoder(config)
	case "console":
		fallthrough
	default:
		return zapcore.NewConsoleEncoder(config)
	}
}
