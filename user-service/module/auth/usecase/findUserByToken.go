package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) FindUserByToken(ctx context.Context, requestId string, req *userRpc.FindUserByTokenRequest) (*userRpc.User, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.FindUserByToken")
	defer span.End()

	claimedToken, err := pkg.ParseClaimFromToken(req.Token, config.Get().JwtAccessTokenSecret)
	if err != nil {
		span.RecordError(err)
		u.logger.Error(fmt.Sprintf("requestId : %s , error parsing token: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, claimedToken.UserID)
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
