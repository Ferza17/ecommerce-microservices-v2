package telemetry

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"net/http"
)

func TelemetryHTTPMiddleware(telemetry telemetryInfrastructure.ITelemetryInfrastructure) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := telemetry.StartSpanFromHttpRequest(r, "Interceptor.TelemetryHTTPMiddleware")
			defer span.End()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
