package user

import (
	"context"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *userService) AuthServiceVerifyIsExcluded(ctx context.Context, requestId string, in *pb.AuthServiceVerifyIsExcludedRequest) (*pb.AuthServiceVerifyIsExcludedResponse, error) {
	md := metadata.New(map[string]string{pkgContext.CtxKeyRequestID: requestId})
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := s.authSvc.AuthServiceVerifyIsExcluded(ctx, in, grpc.Header(&md))
	if err != nil {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}
	return resp, nil
}
