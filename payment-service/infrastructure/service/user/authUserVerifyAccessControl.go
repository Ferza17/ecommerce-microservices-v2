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
	token := pkgContext.GetTokenAuthorizationFromContext(ctx)
	fullMethodName, err := pkgContext.GetFullMethodNameFromContext(ctx)
	if err != nil {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.Error(err))
		return status.Error(codes.Unauthenticated, "missing metadata")
	}

	md := metadata.New(map[string]string{pkgContext.CtxKeyRequestID: token})
	md.Set(pkgContext.CtxKeyAuthorization, fmt.Sprintf("Bearer %s", token))
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		md.Set(key, value)
	}
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := s.authSvc.AuthUserVerifyAccessControl(ctx, &pb.AuthUserVerifyAccessControlRequest{
		Token:          token,
		FullMethodName: &fullMethodName,
	}, grpc.Header(&md))
	if err != nil {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(err))
		return err
	}

	if !resp.IsValid {
		s.logger.Error("UserService.AuthServiceVerifyIsExcluded", zap.String("requestId", requestId), zap.Error(errors.New("UnAuthenticated")))
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}

	return nil
}
