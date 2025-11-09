package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/notification"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.AuthUserLoginByEmailAndPasswordRequest) (*emptypb.Empty, error) {
	var (
		causationId = pkgContext.GetCausationIdFromContext(ctx)
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserLoginByEmailAndPassword")
	defer span.End()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("UserUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	user, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, req.Email, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
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

	if err = u.eventUseCase.AppendEventEnvelope(ctx, &pbEvent.EventEnvelope{
		XId:           primitive.NewObjectID().Hex(),
		EventType:     config.Get().BrokerKafkaTopicNotifications.EmailOtpUserLogin,
		AggregateType: pbEvent.AggregateType_NOTIFICATION,
		AggregateId:   primitive.NewObjectID().Hex(), // Because Notification Service use mongodb
		Version:       0,
		OccurredAt:    timestamppb.New(now),
		CorrelationId: requestId,
		CausationId:   &causationId,
	}, &notificationRpc.SendOtpEmailNotificationRequest{
		Email:            user.Email,
		Otp:              otp,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_LOGIN_OTP,
	}); err != nil {
		u.logger.Error("AuthUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
