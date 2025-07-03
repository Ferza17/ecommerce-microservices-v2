package usecase

import (
	"context"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
)

func (u *authUseCase) AuthUserLogoutByToken(ctx context.Context, requestId string, req *pb.AuthUserLogoutByTokenRequest) (*pb.AuthUserLogoutByTokenResponse, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserLogoutByToken")
	defer span.End()
	//TODO
	// 1. Delete Session related to user
	// 2. Delete Cache related to user
	return nil, nil
}
