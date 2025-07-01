package telemetry

import (
	"context"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TelemetryRPCInterceptor(telemetry telemetryInfrastructure.ITelemetryInfrastructure) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx, span := telemetry.Tracer(ctx, "Interceptor.TelemetryRPCInterceptor")
		defer func() {
			span.End()

			md, _ := metadata.FromIncomingContext(ctx)
			carrier := propagation.MapCarrier{}
			otel.GetTextMapPropagator().Inject(ctx, carrier)
			for key, value := range carrier {
				md.Set(key, value)
			}

			if err == nil {
				// Send metadata as response headers
				if err = grpc.SetHeader(ctx, md); err != nil {
					return
				}
				return
			}
		}()
		return handler(ctx, req)
	}
}
