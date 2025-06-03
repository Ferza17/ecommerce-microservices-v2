package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type (
	IZapLogger interface {
		Info(msg string, fields ...zap.Field)
		Error(msg string)
		Debug(msg string)
		SetFields(fields ...zap.Field) *zap.Logger
	}

	zapLogger struct {
		logger *zap.Logger
	}
)

func NewZapLogger() IZapLogger {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatalf("error when register logger: %v\n", err)
	}
	return &zapLogger{
		logger: logger,
	}
}

func (z *zapLogger) Info(msg string, fields ...zap.Field) {
	z.logger.Info(msg, fields...)
}

func (z *zapLogger) Error(msg string) {
	z.logger.Error(msg)
}

func (z *zapLogger) Debug(msg string) {
	z.logger.Debug(msg)
}

func (z *zapLogger) SetFields(fields ...zap.Field) *zap.Logger {
	return z.logger.With(fields...)
}
