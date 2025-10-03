package usecase

import (
	"context"
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserVerifyOtp(ctx context.Context, requestId string, req *pb.AuthUserVerifyOtpRequest) (*pb.AuthUserVerifyOtpResponse, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserVerifyOtp")
	defer span.End()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	userId, err := u.authRedisRepository.GetOtp(ctx, requestId, req.Otp)
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		if err == redis.Nil {
			return nil, status.Error(codes.NotFound, "otp not found")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if userId == nil {
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(errors.New("userId is nil")))
		return nil, status.Error(codes.NotFound, "not found")
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, *userId, nil)
	if err != nil && err == gorm.ErrRecordNotFound {
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	if user == nil {
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(errors.New("user is nil")))
		return nil, status.Error(codes.NotFound, "not found")
	}

	// Generate Access & Refresh Token
	accessToken, refreshToken, err := u.GenerateAccessToken(ctx, requestId, user)
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error generating token")))
		return nil, err
	}

	user.IsVerified = true
	user.UpdatedAt = &now
	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgUser.Users, user.ID, kafka.JSON_SCHEMA, user); err != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AuthUserVerifyOtpResponse{
		Status:  "success",
		Message: "AuthUserVerifyOtp",
		Data: &pb.AuthUserVerifyOtpResponse_AuthUserVerifyOtpResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
