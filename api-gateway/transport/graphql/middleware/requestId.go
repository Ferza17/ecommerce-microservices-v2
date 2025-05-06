package middleware

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/pkg"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger(logger pkg.IZapLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := middleware.GetReqID(r.Context())
			if reqID == "" {
				reqID = uuid.NewString()
			}

			// Inject request ID ke context logger
			ctx := r.Context()
			ctx = context.WithValue(ctx, enum.XRequestIDHeader.String(), reqID)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				logger.Info("Served",
					zap.String("proto", r.Proto),
					zap.String("path", r.URL.Path),
					zap.Duration("latency", time.Since(t1)),
					zap.Int("status", ww.Status()),
					zap.Int("size", ww.BytesWritten()),
					zap.String("reqId", reqID))
			}()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
