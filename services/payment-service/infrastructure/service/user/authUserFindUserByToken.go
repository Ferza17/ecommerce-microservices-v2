package user

import (
	"context"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *userService) AuthUserFindUserByToken(ctx context.Context, requestId string, in *pb.AuthUserFindUserByTokenRequest) (*pb.AuthUserFindUserByTokenResponse, error) {
	md := metadata.New(map[string]string{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: fmt.Sprintf("Bearer %s", pkgContext.GetTokenAuthorizationFromContext(ctx)),
	})
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}
	resp, err := s.authSvc.AuthUserFindUserByToken(metadata.NewOutgoingContext(ctx, md), &pb.AuthUserFindUserByTokenRequest{
		Token: in.Token,
	}, grpc.Header(&md))
	if err != nil {
		s.logger.Error("UserService.AuthUserFindUserByToken", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}
	return resp, nil
}
