package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *authUseCase) UserLoginByEmailAndPassword(ctx context.Context, requestId string, req *pb.UserLoginByEmailAndPasswordRequest) error {
	var (
		err                      error
		reqUserLoginNotification *pb.SendOtpEmailNotificationRequest
		tx                       = u.userPostgresqlRepository.OpenTransactionWithContext(ctx)
		eventStore               = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.UserService.String(),
			EventType:     enum.USER_LOGIN.String(),
			Status:        enum.SUCCESS.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserLoginByEmailAndPassword")

	defer func(err error, eventStore *pb.EventStore) {
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
			eventStore.Status = enum.FAILED.String()
		}

		// Success Login
		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}

		loginNotificationPayload, err := util.ConvertStructToProtoStruct(reqUserLoginNotification)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		}

		eventStore.Payload = loginNotificationPayload
		eventStore.EventType = enum.NOTIFICATION_EMAIL_OTP.String()
		eventStore.Service = enum.NotificationService.String()
		eventStore.Status = enum.PENDING.String()

		eventStoreMessage, err = proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		}

		// PENDING Notification Login
		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
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

	reqUserLoginNotification = &pb.SendOtpEmailNotificationRequest{
		Email:            user.Email,
		Otp:              otp,
		NotificationType: pb.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP,
	}

	message, err := proto.Marshal(reqUserLoginNotification)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return status.Error(codes.Internal, err.Error())
	}

	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.NotificationExchange, enum.NOTIFICATION_EMAIL_OTP, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publish message err : %v", err))
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
