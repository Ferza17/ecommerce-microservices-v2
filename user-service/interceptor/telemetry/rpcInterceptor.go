package telemetry

import (
	"context"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"google.golang.org/grpc"
)

func TelemetryRPCInterceptor(telemetry telemetryInfrastructure.ITelemetryInfrastructure) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx, span := telemetry.Tracer(ctx, "Interceptor.TelemetryRPCInterceptor")
		defer span.End()
		return handler(ctx, req)
	}
}
