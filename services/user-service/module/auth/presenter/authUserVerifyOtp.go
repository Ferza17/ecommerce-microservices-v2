package presenter

import (
	"context"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *AuthPresenter) AuthUserVerifyOtp(ctx context.Context, req *userRpc.AuthUserVerifyOtpRequest) (*userRpc.AuthUserVerifyOtpResponse, error) {
	ctx, span := p.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthPresenter.AuthUserVerifyOtp")
	defer span.End()
	requestID := pkgContext.GetRequestIDFromContext(ctx)

	if err := req.Validate(); err != nil {
		p.logger.Error("AuthPresenter.AuthUserVerifyOtp", zap.String("requestID", requestID), zap.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := p.authUseCase.AuthUserVerifyOtp(ctx, requestID, req)
	if err != nil {
		p.logger.Error("AuthPresenter.AuthUserVerifyOtp", zap.String("requestID", requestID), zap.Error(err))
		return nil, err
	}

	return resp, nil
}
