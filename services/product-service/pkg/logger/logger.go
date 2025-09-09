package logger

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type (
	IZapLogger interface {
		Info(fnName string, f ...zap.Field)
		Error(fnName string, f ...zap.Field)
		Debug(fnName string, f ...zap.Field)
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
	return &zapLogger{
		logger: logger,
	}
}

func (z *zapLogger) Info(fnName string, fields ...zap.Field) {
	fields = append(fields, zap.String("service", config.Get().ProductServiceServiceName))
	fields = append(fields, zap.String("context", fnName))
	z.logger.Info(fnName, fields...)
}

func (z *zapLogger) Error(fnName string, fields ...zap.Field) {
	fields = append(fields, zap.String("service", config.Get().ProductServiceServiceName))
	fields = append(fields, zap.String("context", fnName))
	z.logger.Error(fnName, fields...)
}

func (z *zapLogger) Debug(fnName string, fields ...zap.Field) {
	fields = append(fields, zap.String("service", config.Get().ProductServiceServiceName))
	fields = append(fields, zap.String("msg", fnName))
	z.logger.Debug(fnName, fields...)
}
