package logger

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/cqroot/garden/internal/config"
)

type Logger struct {
	zapLogger *zap.Logger
}

func getZapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006/01/02-15:04:05"))
	})
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return encoder
}

func New(config *config.Config) *Logger {
	logger := Logger{}

	logLevel := strings.ToLower(config.LogLevel())

	writeSyncer := zapcore.AddSync(os.Stdout)

	var core zapcore.Core

	switch logLevel {
	case "debug":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.DebugLevel)
	case "info":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.InfoLevel)
	case "warn":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.WarnLevel)
	case "error":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.ErrorLevel)
	}

	if config.LogWithCaller() {
		logger.zapLogger = zap.New(core, zap.AddCaller())
	} else {
		logger.zapLogger = zap.New(core)
	}

	if logLevel == "debug" {
		logger.Debug("Enable debug output")
	}

	return &logger
}

func (logger *Logger) Debug(msg string, fields ...zap.Field) {
	logger.zapLogger.Debug(msg, fields...)
}

func (logger *Logger) Info(msg string, fields ...zap.Field) {
	logger.zapLogger.Info(msg, fields...)
}

func (logger *Logger) Error(msg string, fields ...zap.Field) {
	logger.zapLogger.Error(msg, fields...)
}

func (logger *Logger) Fatal(msg string, fields ...zap.Field) {
	logger.zapLogger.Fatal(msg, fields...)
}
