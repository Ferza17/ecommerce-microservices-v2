package presenter

import (
	"context"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *AuthPresenter) AuthUserLoginByEmailAndPassword(ctx context.Context, req *pb.AuthUserLoginByEmailAndPasswordRequest) (*pb.AuthUserLoginByEmailAndPasswordResponse, error) {
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "AuthPresenter.AuthUserLoginByEmailAndPassword")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := req.Validate(); err != nil {
		p.logger.Error("AuthPresenter.AuthUserLoginByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := p.authUseCase.AuthUserLoginByEmailAndPassword(ctx, requestID, req)
	if err != nil {
		p.logger.Error("AuthPresenter.AuthUserLoginByEmailAndPassword", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	return resp, nil
}
