package metric

import (
	"fmt"
	"net/http"
	"time"

	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/metric"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/response"
	"go.opentelemetry.io/otel/attribute"
)

func MetricHTTPMiddleware(telemetry telemetryInfrastructure.ITelemetryInfrastructure) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/metrics" {
				next.ServeHTTP(w, r)
				return
			}
			start := time.Now()
			ctx, span := telemetry.StartSpanFromHttpRequest(r, "Interceptor.MetricHTTPMiddleware")
			defer span.End()

			// Wrap the response writer to capture status code
			ww := &response.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
			r = r.WithContext(ctx)

			next.ServeHTTP(ww, r)

			duration := time.Since(start).Seconds()
			status := fmt.Sprintf("%d", ww.StatusCode)

			pkgMetric.HttpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
			pkgMetric.HttpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
			span.SetAttributes(attribute.Int("http.status_code", ww.StatusCode))
		})
	}
}
