package logger

import (
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type (
	IZapLogger interface {
		Info(msg string)
		Error(msg string)
		Debug(msg string)
	}

	zapLogger struct {
		logger *zap.Logger
	}
)

var Set = wire.NewSet(NewZapLogger)

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
	log.Println("LOGGER registered")
	return &zapLogger{
		logger: logger,
	}
}

func (z *zapLogger) Info(msg string) {
	fields := []zap.Field{
		zap.String("service", config.Get().ServiceName),
		zap.String("msg", msg),
	}
	z.logger.Info(msg, fields...)
}

func (z *zapLogger) Error(msg string) {
	fields := []zap.Field{
		zap.String("service", config.Get().ServiceName),
		zap.String("msg", msg),
		zap.Error(errors.New(msg)),
	}
	z.logger.Error("error", fields...)
}

func (z *zapLogger) Debug(msg string) {
	fields := []zap.Field{
		zap.String("service", config.Get().ServiceName),
		zap.String("msg", msg),
	}
	z.logger.Debug(msg, fields...)
}
