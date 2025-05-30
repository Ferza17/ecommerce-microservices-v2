package usecase

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (u *authUseCase) UserVerifyOtp(ctx context.Context, requestId string, req *userRpc.UserVerifyOtpRequest) (*userRpc.UserVerifyOtpResponse, error) {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserVerifyOtp")
	defer span.End()

	userId, err := u.authRedisRepository.GetOtp(ctx, requestId, req.Otp)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		return nil, err
	}

	if userId == nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error getting otp: %v", requestId, err))
		return nil, status.Error(codes.NotFound, "not found")
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, *userId)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by id: %v", requestId, err))
		return nil, status.Error(codes.NotFound, "not found")
	}

	var (
		te  = config.Get().JwtAccessTokenExpirationTime
		now = time.Now().UTC()
	)

	accessToken, err := pkg.GenerateToken(pkg.Claim{
		UserID:    user.ID,
		CreatedAt: &now,
		StandardClaims: jwt.StandardClaims{
			Audience:  enum.UserService.String(),
			ExpiresAt: now.Add(te).Unix(),
		},
	}, config.Get().JwtAccessTokenSecret)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error generating token: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := pkg.GenerateToken(pkg.Claim{
		UserID:    user.ID,
		CreatedAt: &now,
		StandardClaims: jwt.StandardClaims{
			Audience:  enum.UserService.String(),
			ExpiresAt: now.Add(config.Get().JwtRefreshTokenExpirationTime).Unix(),
		},
	}, config.Get().JwtRefreshTokenSecret)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error generating refresh token: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userRpc.UserVerifyOtpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
