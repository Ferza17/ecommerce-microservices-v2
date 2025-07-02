package metric

import (
	"fmt"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/response"
	"net/http"
	"time"
)

func MetricHTTPMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/metrics" {
				next.ServeHTTP(w, r)
				return
			}

			start := time.Now()

			// Wrap the response writer to capture status code
			ww := &response.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}

			next.ServeHTTP(ww, r)

			duration := time.Since(start).Seconds()

			pkgMetric.HttpRequests.WithLabelValues(r.Method, r.URL.Path, fmt.Sprintf("%d", ww.StatusCode)).Inc()
			pkgMetric.HttpDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		})
	}
}
