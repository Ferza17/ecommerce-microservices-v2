package logger

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"net/http"
	"time"
)

func LoggerHTTPMiddleware(
	logger logger.IZapLogger,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logger.Info(fmt.Sprintf("%s %s", r.URL, time.Now().Sub(startTime)))

			defer func() {
				logger.Info(fmt.Sprintf("%s - %s", r.URL, time.Since(startTime)))
			}()
			next.ServeHTTP(w, r)
		})
	}
}
