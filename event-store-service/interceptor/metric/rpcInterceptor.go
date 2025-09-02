package metric

import (
	"context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/metric"
	"google.golang.org/grpc"
	"time"
)

func MetricUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start).Seconds()
		status := "success"
		if err != nil {
			status = "error"
		}

		pkgMetric.GrpcRequestsTotal.WithLabelValues(info.FullMethod, status).Inc()
		pkgMetric.GrpcRequestDuration.WithLabelValues(info.FullMethod).Observe(duration)

		return resp, err
	}
}
