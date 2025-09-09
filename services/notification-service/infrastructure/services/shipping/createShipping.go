package shipping

import (
	"context"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/shipping"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *shippingService) CreateShipping(ctx context.Context, requestId string, request *pb.CreateShippingRequest) (*pb.CreateShippingResponse, error) {
	md := metadata.New(map[string]string{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: fmt.Sprintf("Bearer %s", pkgContext.GetTokenAuthorizationFromContext(ctx)),
	})
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}

	resp, err := s.shippingSvc.CreateShipping(metadata.NewOutgoingContext(ctx, md), request, grpc.Header(&md))
	if err != nil {
		s.logger.Error("ShippingService.CreateShipping", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	return resp, nil
}
