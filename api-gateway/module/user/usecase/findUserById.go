package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"
)

func (u *UserUseCase) FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindUserById")
	defer span.End()

	user, err := u.userService.FindUserById(ctx, requestId, req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding user : %v", err))
		return nil, err
	}
	return user, nil
}
