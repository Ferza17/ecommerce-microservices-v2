package usecase

import (
	"context"
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) GenerateAccessToken(ctx context.Context, requestId string, user *orm.User) (accessToken string, refreshToken string, err error) {
	// Generate Access & Refresh Token
	var (
		defaultAccessTokenCfg  = token.DefaultAccessTokenConfig()
		defaultRefreshTokenCfg = token.DefaultRefreshTokenConfig()

		accessControls []*pb.AccessControl
	)

	if user.Role != nil && len(user.Role.AccessControls) > 0 {
		accessControls = orm.AccessControlsToProto(user.Role.AccessControls)
	}

	accessToken, err = token.GenerateToken(
		token.GenerateClaim(user.ToProto(), user.Role.ToProto(), accessControls, defaultAccessTokenCfg),
		defaultAccessTokenCfg,
	)
	if err != nil {
		u.logger.Error("AuthUseCase.GenerateAccessToken", zap.String("requestId", requestId), zap.Error(errors.New("error generating access token")))
		return "", "", status.Error(codes.Internal, err.Error())
	}

	refreshToken, err = token.GenerateToken(
		token.GenerateClaim(user.ToProto(), user.Role.ToProto(), accessControls, defaultRefreshTokenCfg),
		defaultRefreshTokenCfg,
	)
	if err != nil {
		u.logger.Error("AuthUseCase.GenerateAccessToken", zap.String("requestId", requestId), zap.Error(errors.New("error generating refresh token")))
		return "", "", status.Error(codes.Internal, err.Error())
	}

	return accessToken, refreshToken, nil
}
