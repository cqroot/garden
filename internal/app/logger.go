package app

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
)

func getZapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006/01/02-15:04:05"))
	})
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return encoder
}

func InitLogger(logLevel string, caller bool) {
	writeSyncer := zapcore.AddSync(os.Stdout)

	var core zapcore.Core

	level := strings.ToLower(logLevel)
	switch level {
	case "debug":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.DebugLevel)
	case "info":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.InfoLevel)
	case "warn":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.WarnLevel)
	case "error":
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.ErrorLevel)
	}

	if caller {
		zapLogger = zap.New(core, zap.AddCaller())
	} else {
		zapLogger = zap.New(core)
	}
}

func Logger() *zap.Logger {
	return zapLogger
}
