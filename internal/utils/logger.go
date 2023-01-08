package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
)

func getZapEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return encoder
}

func InitLogger(debug bool, caller bool) {
	writeSyncer := zapcore.AddSync(os.Stdout)

	var core zapcore.Core
	if debug {
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.DebugLevel)
	} else {
		core = zapcore.NewCore(getZapEncoder(), writeSyncer, zapcore.InfoLevel)
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
