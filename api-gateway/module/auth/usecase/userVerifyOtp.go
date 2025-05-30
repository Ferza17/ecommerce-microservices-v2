package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"
)

func (u *authUseCase) UserVerifyOtp(ctx context.Context, requestId string, req *userRpc.UserVerifyOtpRequest) (*userRpc.UserVerifyOtpResponse, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserLoginByEmailAndPassword")
	defer span.End()

	resp, err := u.authService.UserVerifyOtp(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error verifying otp: %s", err.Error()))
		return nil, err
	}

	return resp, nil
}
