package usecase

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/notification"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func (u *authUseCase) SentOTP(ctx context.Context, requestId string, user *pb.User) error {
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "AuthUseCase.SentOTP")
	defer span.End()

	otp := util.GenerateOTP()
	if err := u.authRedisRepository.SetOtp(ctx, requestId, otp, user.Id); err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	reqUserEmailOtp := &notificationRpc.SendOtpEmailNotificationRequest{
		Email:            user.Email,
		Otp:              otp,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP,
	}

	message, err := proto.Marshal(reqUserEmailOtp)
	if err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeNotification, config.Get().QueueNotificationEmailOtpCreated, message); err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
