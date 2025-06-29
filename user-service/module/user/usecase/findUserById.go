package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
)

func (u *userUseCase) FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	tx := u.postgresSQLInfrastructure.GormDB.Begin()
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindUserById")
	defer span.End()
	fetchedUser, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id, tx)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, err
	}
	tx.Commit()
	return fetchedUser.ToProto(), nil
}
