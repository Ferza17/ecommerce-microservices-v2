package telemetry

import (
	"context"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TelemetryRPCInterceptor(telemetry telemetryInfrastructure.ITelemetryInfrastructure) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx, span := telemetry.StartSpanFromRpcMetadata(ctx, "Interceptor.TelemetryRPCInterceptor")

		span.SetAttributes(
			attribute.String("rpc.system", "grpc"),
			attribute.String("rpc.method", info.FullMethod),
		)

		defer func() {
			span.End()

			md, _ := metadata.FromIncomingContext(ctx)
			carrier := telemetry.InjectSpanToTextMapPropagator(ctx)
			for key, value := range carrier {
				md.Set(key, value)
			}

			if err == nil {
				if err = grpc.SetHeader(ctx, md); err != nil {
					return
				}
				return
			}
		}()
		return handler(ctx, req)
	}
}
