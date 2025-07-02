package telemetry

import (
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func TelemetryHTTPMiddleware(telemetry telemetryInfrastructure.ITelemetryInfrastructure) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, span := telemetry.Tracer(ctx, "Interceptor.TelemetryHTTPMiddleware")
			defer func() {
				span.End()

				md, _ := metadata.FromIncomingContext(ctx)
				if md != nil {
					carrier := propagation.MapCarrier{}
					otel.GetTextMapPropagator().Inject(ctx, carrier)
					for key, value := range carrier {
						md.Set(key, value)
					}
				}

			}()

			r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
