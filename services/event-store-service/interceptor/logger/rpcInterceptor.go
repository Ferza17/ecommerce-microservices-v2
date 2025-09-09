package logger

import (
	"context"
	"time"

	pkgContext "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func LoggerRPCInterceptor(logger logger.IZapLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		start := time.Now()
		defer func() {
			end := time.Now()
			logger.Info(
				"Interceptor.LoggerRPCInterceptor",
				zap.String("fullMethod", info.FullMethod),
				zap.String("requestId", pkgContext.GetRequestIDFromContext(ctx)),
				zap.Time("start", start),
				zap.Any("request", req),
				zap.Any("response", resp),
				zap.Duration("duration", time.Since(start)),
				zap.Time("end", end),
			)
		}()
		return handler(ctx, req)
	}
}
