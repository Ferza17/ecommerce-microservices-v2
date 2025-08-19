package usecase

import (
	"context"
	"errors"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *authUseCase) AuthUserVerifyOtp(ctx context.Context, requestId string, req *pb.AuthUserVerifyOtpRequest) (*pb.AuthUserVerifyOtpResponse, error) {
	tx := u.postgresSQL.GormDB().Begin()
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserVerifyOtp")
	defer span.End()

	userId, err := u.authRedisRepository.GetOtp(ctx, requestId, req.Otp)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if userId == nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(errors.New("userId is nil")))
		return nil, status.Error(codes.NotFound, "not found")
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, *userId, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(errors.New("user is nil")))
		return nil, status.Error(codes.NotFound, "not found")
	}

	// Generate Access & Refresh Token
	accessToken, refreshToken, err := u.GenerateAccessToken(ctx, requestId, user)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error generating token")))
		return nil, err
	}

	tx.Commit()
	return &pb.AuthUserVerifyOtpResponse{
		Status:  "success",
		Message: "AuthUserVerifyOtp",
		Data: &pb.AuthUserVerifyOtpResponse_AuthUserVerifyOtpResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
