package usecase

import (
	"context"
	"errors"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *userUseCase) FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UserUseCase.FindUserById")
	defer span.End()

	tx := u.postgresSQLInfrastructure.GormDB.Begin()
	fetchedUser, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error("UserUseCase.FindUserById", zap.String("requestId", requestId), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	tx.Commit()
	return fetchedUser.ToProto(), nil
}
