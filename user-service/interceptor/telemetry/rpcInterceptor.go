package telemetry

import (
	"context"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TelemetryRPCInterceptor(telemetry telemetryInfrastructure.ITelemetryInfrastructure) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx, span := telemetry.Tracer(ctx, "Interceptor.TelemetryRPCInterceptor")
		traceID := span.SpanContext().TraceID().String()
		// Define response header metadata
		header := metadata.Pairs(
			"x-trace-id", traceID,
		)
		defer func() {
			span.End()

			if err == nil {
				// Send metadata as response headers
				if err = grpc.SetHeader(ctx, header); err != nil {
					return
				}
				return
			}
		}()
		return handler(ctx, req)
	}
}
