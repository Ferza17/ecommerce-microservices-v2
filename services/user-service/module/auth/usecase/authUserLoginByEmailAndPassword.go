package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/notification"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.AuthUserLoginByEmailAndPasswordRequest) (*emptypb.Empty, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserLoginByEmailAndPassword")
	defer span.End()

	user, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, req.Email, nil)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("user not found")))
			return nil, status.Error(codes.NotFound, err.Error())
		}

		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email : %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	reqHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error hashing password")))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword(reqHashedPassword, []byte(req.Password)); err != nil {
		u.logger.Error("AuthUseCase.AuthUserLoginByEmailAndPassword", zap.String("requestId", requestId), zap.Error(errors.New("error comparing password")))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Send OTP
	otp := util.GenerateOTP()
	if err = u.authRedisRepository.SetOtp(ctx, requestId, otp, user.ID); err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	message, err := proto.Marshal(&notificationRpc.SendOtpEmailNotificationRequest{
		Email:            user.Email,
		Otp:              otp,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_LOGIN_OTP,
	})
	if err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = u.kafkaInfrastructure.Publish(ctx, config.Get().BrokerKafkaTopicNotifications.EmailOtpUserRegister, requestId, message); err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
