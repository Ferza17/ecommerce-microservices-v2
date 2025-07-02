package metric

import (
	"context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/metric"
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

		pkgMetric.GrpcRequests.WithLabelValues(info.FullMethod, status).Inc()
		pkgMetric.GrpcDuration.WithLabelValues(info.FullMethod).Observe(duration)

		return resp, err
	}
}
