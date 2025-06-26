package usecase

import (
	"context"
	"fmt"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) UserVerifyOtp(ctx context.Context, requestId string, req *userRpc.AuthVerifyOtpRequest) (*userRpc.AuthVerifyOtpResponse, error) {
	tx := u.postgresSQL.GormDB.Begin()
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserVerifyOtp")
	defer span.End()

	userId, err := u.authRedisRepository.GetOtp(ctx, requestId, req.Otp)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		return nil, err
	}

	if userId == nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		return nil, status.Error(codes.NotFound, "not found")
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, *userId, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, status.Error(codes.NotFound, "not found")
	}

	var (
		defaultAccessTokenCfg  = token.DefaultAccessTokenConfig()
		defaultRefreshTokenCfg = token.DefaultRefreshTokenConfig()

		accessControls = []*userRpc.AccessControl{}
	)

	if user.Role != nil && len(user.Role.AccessControls) > 0 {
		for _, control := range user.Role.AccessControls {
			accessControls = append(accessControls, control.ToProto())
		}
	}

	accessToken, err := token.GenerateToken(
		token.GenerateClaim(user.ToProto(), user.Role.ToProto(), accessControls, defaultAccessTokenCfg),
		defaultAccessTokenCfg,
	)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error generating accessToken: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := token.GenerateToken(
		token.GenerateClaim(user.ToProto(), user.Role.ToProto(), accessControls, defaultRefreshTokenCfg),
		defaultRefreshTokenCfg,
	)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error refreshToken token: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	tx.Commit()
	return &userRpc.AuthVerifyOtpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
