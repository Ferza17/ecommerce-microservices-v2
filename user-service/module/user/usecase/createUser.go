package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (u *userUseCase) CreateUser(ctx context.Context, requestId string, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	var (
		err             error
		reqUserEmailOtp *pb.SendOtpEmailNotificationRequest
		tx              = u.userPostgresqlRepository.OpenTransactionWithContext(ctx)
		now             = time.Now().UTC()
		eventStore      = &pb.EventStore{
			RequestId:     requestId,
			Service:       enum.UserService.String(),
			EventType:     enum.USER_CREATED.String(),
			Status:        enum.SUCCESS.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.CreateUser")

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
			eventStore.Status = enum.FAILED.String()
		}

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
		return nil, err
	}
	req.Password = string(hashedPassword)

	result, err := u.userPostgresqlRepository.CreateUserWithTransaction(ctx, requestId, &orm.User{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		CreatedAt:   &now,
		UpdatedAt:   &now,
		DiscardedAt: nil,
	}, tx)

	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error creating user: %v", requestId, err))
		return nil, err
	}

	otp := util.GenerateOTP()
	if err = u.authRedisRepository.SetOtp(ctx, requestId, otp, result); err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error setting otp: %v", requestId, err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	reqUserEmailOtp = &pb.SendOtpEmailNotificationRequest{
		Email:            req.Email,
		Otp:              otp,
		NotificationType: pb.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP,
	}

	message, err := proto.Marshal(reqUserEmailOtp)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.NotificationExchange, enum.NOTIFICATION_EMAIL_OTP, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publish message err : %v", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	tx.Commit()
	return &pb.CreateUserResponse{
		Id: result,
	}, nil
}
