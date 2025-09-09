package logger

import (
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func LoggerHTTPMiddleware(
	logger logger.IZapLogger,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			start := time.Now()
			defer func() {
				end := time.Now()
				logger.Info(
					"Interceptor.LoggerHTTPMiddleware",
					zap.String("http_url", r.URL.String()),
					zap.String("requestId", pkgContext.GetRequestIDFromContext(r.Context())),
					zap.Time("start", start),
					//zap.Any("request", r),
					zap.Duration("duration", time.Since(start)),
					zap.Time("end", end),
				)
			}()
			r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
