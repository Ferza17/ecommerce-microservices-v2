package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) FindUserByToken(ctx context.Context, requestId string, req *userRpc.AuthFindUserByTokenRequest) (*userRpc.User, error) {
	tx := u.postgresSQL.GormDB.Begin()
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindUserByToken")
	defer span.End()

	claimedToken, err := token.ValidateJWTToken(req.Token, token.DefaultRefreshTokenConfig())
	if err != nil {
		tx.Rollback()
		span.RecordError(err)
		u.logger.Error(fmt.Sprintf("requestId : %s , error parsing token: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, claimedToken.GetUserID(), tx)
	if err != nil {
		span.RecordError(err)
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userRpc.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
