package usecase

import (
	"context"
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) GenerateAccessToken(ctx context.Context, requestId string, user *orm.User) (accessToken string, refreshToken string, err error) {
	var (
		defaultAccessTokenCfg  = token.DefaultAccessTokenConfig()
		defaultRefreshTokenCfg = token.DefaultRefreshTokenConfig()
	)

	accessToken, err = token.GenerateToken(
		token.GenerateClaim(user.ToProto(), defaultAccessTokenCfg),
		defaultAccessTokenCfg,
	)
	if err != nil {
		u.logger.Error("AuthUseCase.GenerateAccessToken", zap.String("requestId", requestId), zap.Error(errors.New("error generating access token")))
		return "", "", status.Error(codes.Internal, err.Error())
	}

	refreshToken, err = token.GenerateToken(
		token.GenerateClaim(user.ToProto(), defaultRefreshTokenCfg),
		defaultRefreshTokenCfg,
	)

	if err != nil {
		u.logger.Error("AuthUseCase.GenerateAccessToken", zap.String("requestId", requestId), zap.Error(errors.New("error generating refresh token")))
		return "", "", status.Error(codes.Internal, err.Error())
	}

	return accessToken, refreshToken, nil
}
