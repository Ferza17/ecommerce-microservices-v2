package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"
)

func (u *userUseCase) FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindUserById")
	defer span.End()
	fetchedUser, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, err
	}
	return &userRpc.User{
		Id:    fetchedUser.ID,
		Name:  fetchedUser.Name,
		Email: fetchedUser.Email,
	}, nil
}
