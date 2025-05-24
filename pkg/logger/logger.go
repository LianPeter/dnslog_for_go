package logger

import (
	"dnslog_for_go/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var log *zap.Logger

// InitLogger 初始化日志系统
func InitLogger(cfg *config.LogConfig) error {
	// 配置日志输出
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,    // 每个日志文件最大MB
		MaxBackups: cfg.MaxBackups, // 保留的旧日志文件数量
		MaxAge:     cfg.MaxAge,     // 保留的最大天数
		Compress:   cfg.Compress,   // 是否压缩
	})

	// 设置日志级别
	var level zapcore.Level
	err := level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return err
	}

	// 创建编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		level,
	)

	// 创建日志记录器
	log = zap.New(core)
	zap.ReplaceGlobals(log)

	return nil
}

// Info 记录信息级别的日志
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Error 记录错误级别的日志
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Fatal 记录致命错误并退出程序
func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

// Debug 记录调试级别的日志
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Warn 记录警告级别的日志
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}
