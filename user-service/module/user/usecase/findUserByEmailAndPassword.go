package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUseCase) FindUserByEmailAndPassword(ctx context.Context, requestId string, request *userRpc.FindUserByEmailAndPasswordRequest) (*userRpc.User, error) {
	var (
		tx = u.postgresSQLInfrastructure.GormDB.Begin()
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindUserByEmailAndPassword")
	defer span.End()

	user, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, request.Email, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email and password: %v", requestId, err))
		return nil, err
	}

	reqHashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword(reqHashedPassword, []byte(request.Password)); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error comparing password: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &userRpc.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil

}
