package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/event/v1"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/notification/v1"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *authUseCase) UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *userRpc.UserLoginByEmailAndPasswordRequest) error {
	var (
		err                      error
		reqUserLoginNotification *notificationRpc.SendOtpEmailNotificationRequest
		tx                       = u.userPostgresqlRepository.OpenTransactionWithContext(ctx)
		eventStore               = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ServiceName,
			EventType:     config.Get().QueueUserLogin,
			Status:        config.Get().CommonSagaStatusSuccess,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserLoginByEmailAndPassword")

	defer func(err error, eventStore *eventRpc.EventStore) {
		defer span.End()
		payload, err := util.ConvertStructToProtoStruct(req)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		}
		eventStore.Payload = payload

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		}

		if err != nil {
			tx.Rollback()
			eventStore.Status = config.Get().CommonSagaStatusFailed
		}

		// Success Login
		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}

		loginNotificationPayload, err := util.ConvertStructToProtoStruct(reqUserLoginNotification)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		}

		eventStore.Payload = loginNotificationPayload
		eventStore.EventType = config.Get().QueueNotificationEmailOtpCreated
		eventStore.Service = config.Get().NotificationServiceName
		eventStore.Status = config.Get().CommonSagaStatusPending

		eventStoreMessage, err = proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		}

		// PENDING Notification Login
		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}

		tx.Commit()
	}(err, eventStore)

	user, err := u.userPostgresqlRepository.FindUserByEmailWithTransaction(ctx, requestId, req.Email, tx)
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error finding user by email and password: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	reqHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword(reqHashedPassword, []byte(req.Password)); err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error comparing password: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	otp := util.GenerateOTP()
	if err = u.authRedisRepository.SetOtp(ctx, requestId, otp, user.ID); err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error setting otp: %v", requestId, err))
		return status.Error(codes.Internal, err.Error())
	}

	reqUserLoginNotification = &notificationRpc.SendOtpEmailNotificationRequest{
		Email:            user.Email,
		Otp:              otp,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP,
	}

	message, err := proto.Marshal(reqUserLoginNotification)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return status.Error(codes.Internal, err.Error())
	}

	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeNotification, config.Get().QueueNotificationEmailOtpCreated, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publish message err : %v", err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
