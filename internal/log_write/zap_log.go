package log_write

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	Zap  *zap.Logger
	once sync.Once
)

func InitZapLogger() {
	once.Do(func() {
		data := time.Now().Format("2006-01-02")

		// 项目根下 logs 目录
		projectDir, _ := os.Getwd()
		logDir := filepath.Join(projectDir, "logs")
		_ = os.MkdirAll(logDir, 0755)

		logFile := filepath.Join(logDir, data+".log")

		// 自定义 Encoder（输出格式）
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		// 设置日志输出文件和级别
		fileWriter, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(fileWriter),
			zapcore.DebugLevel,
		)

		Zap = zap.New(core, zap.AddCaller())
	})
}

/*封装几个类型的日志信息，方便调用*/

func Info(msg string, fields ...zap.Field) {
	if Zap != nil {
		Zap.Info(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if Zap != nil {
		Zap.Debug(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Zap != nil {
		Zap.Warn(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if Zap != nil {
		Zap.Error(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if Zap != nil {
		Zap.Fatal(msg, fields...)
	}
}
