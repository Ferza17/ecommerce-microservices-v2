package user

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	pb "github.com/ferza17/ecommerce-microservices-v2/product-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *userService) AuthUserVerifyAccessControl(ctx context.Context, requestId string, in *pb.AuthUserVerifyAccessControlRequest) (*pb.AuthUserVerifyAccessControlResponse, error) {
	md := metadata.New(map[string]string{enum.XRequestIDHeader.String(): requestId})
	md.Set(pkgContext.CtxKeyAuthorization, fmt.Sprintf("Bearer %s", in.GetToken()))
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := s.authSvc.AuthUserVerifyAccessControl(ctx, in, grpc.Header(&md))
	if err != nil {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}
	return resp, nil
}
