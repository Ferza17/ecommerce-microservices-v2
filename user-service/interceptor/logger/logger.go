package logger

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

//TODO:
//1. Need to enhance logger PKG
//2. Need to enhance middleware logger

func LoggerRPCInterceptor(logger logger.IZapLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		startTime := time.Now()
		logger.Info(fmt.Sprintf("%s %s", info.FullMethod, time.Now().Sub(startTime)))

		resp, err = handler(ctx, req)

		duration := time.Since(startTime)
		logger.Info(fmt.Sprintf("%s - %s", info.FullMethod, duration))

		return resp, err
	}
}

func LoggerHTTPMiddleware(
	logger logger.IZapLogger,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logger.Info(fmt.Sprintf("%s %s", r.URL, time.Now().Sub(startTime)))

			next.ServeHTTP(w, r)

			duration := time.Since(startTime)
			logger.Info(fmt.Sprintf("%s - %s", r.URL, duration))
		})
	}
}
