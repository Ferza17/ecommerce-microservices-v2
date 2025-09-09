package user

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *userService) AuthUserVerifyAccessControl(ctx context.Context, requestId string) error {
	fullMethodName, err := pkgContext.GetFullMethodNameFromContext(ctx)
	if err != nil {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.Error(err))
		return status.Error(codes.Unauthenticated, "missing metadata")
	}
	md := metadata.New(map[string]string{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: fmt.Sprintf("Bearer %s", pkgContext.GetTokenAuthorizationFromContext(ctx)),
	})
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}
	resp, err := s.authSvc.AuthUserVerifyAccessControl(metadata.NewOutgoingContext(ctx, md), &pb.AuthUserVerifyAccessControlRequest{
		Token:          pkgContext.GetTokenAuthorizationFromContext(ctx),
		FullMethodName: &fullMethodName,
	}, grpc.Header(&md))
	if err != nil {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	if !resp.Data.IsValid {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(errors.New("UnAuthenticated")))
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}

	return nil
}
