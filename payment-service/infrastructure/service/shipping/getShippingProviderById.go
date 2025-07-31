package shipping

import (
	"context"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/shipping"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *shippingService) GetShippingProviderById(ctx context.Context, requestId string, request *pb.GetShippingProviderByIdRequest) (*pb.GetShippingProviderByIdResponse, error) {
	md := metadata.New(map[string]string{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: fmt.Sprintf("Bearer %s", pkgContext.GetTokenAuthorizationFromContext(ctx)),
	})
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}
	resp, err := s.shippingProviderSvc.GetShippingProviderById(metadata.NewOutgoingContext(ctx, md), request, grpc.Header(&md))
	if err != nil {
		s.logger.Error("ShippingService.GetShippingProviderById", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	return resp, nil
}
