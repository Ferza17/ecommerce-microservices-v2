package usecase

import (
	"context"
	"errors"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/notification"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (u *userUseCase) CreateUser(ctx context.Context, requestId string, req *pb.AuthUserRegisterRequest) (*pb.AuthUserRegisterResponse, error) {
	var (
		err         error
		causationId = pkgContext.GetCausationIdFromContext(ctx)
	)

	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "UserUseCase.CreateUser")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("UserUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// Validate is email already exists
	existedUser, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, req.Email, nil)
	if err != nil && err != gorm.ErrRecordNotFound {
		u.logger.Error("UserUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if existedUser != nil {
		u.logger.Error("UserUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("user with this email already exists")))
		return nil, status.Error(codes.AlreadyExists, "User with this email already exists")
	}

	// Validate Role
	role, err := u.rolePostgresqlRepository.FindRoleByName(ctx, requestId, req.Role.String(), nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.logger.Error("UserUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("role not found")))
			return nil, status.Error(codes.NotFound, "role not found")
		}
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("UserUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("error hashing password")))
		return nil, err
	}
	req.Password = string(hashedPassword)
	user := orm.UserFromProto(&pb.User{
		Id:         uuid.NewString(),
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		IsVerified: false,
		Role:       role.ToProto(),
		CreatedAt:  timestamppb.New(now),
		UpdatedAt:  timestamppb.New(now),
	})

	// Sent OTP
	otp := util.GenerateOTP()
	if err = u.authRedisRepository.SetOtp(ctx, requestId, otp, user.ID); err != nil {
		u.logger.Error("UserUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// SEND TO SINK CONNECTOR PG USER
	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgUser.Users, user.ID, kafka.JSON_SCHEMA, user); err != nil {
		u.logger.Error("UserUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, err
	}

	// SEND TO OUTBOX
	payload, err := util.ProtobufToBase64(&notificationRpc.SendOtpEmailNotificationRequest{
		Email:            user.Email,
		Otp:              otp,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_REGISTER_OTP,
	})
	if err != nil {
		u.logger.Error("UserUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = u.eventUseCase.AppendEventEnvelope(ctx, &pbEvent.EventEnvelope{
		XId:           primitive.NewObjectID().Hex(),
		EventType:     config.Get().BrokerKafkaTopicNotifications.EmailOtpUserRegister,
		AggregateType: pbEvent.AggregateType_NOTIFICATION,
		AggregateId:   primitive.NewObjectID().Hex(), // Because Notification Service use mongodb
		Version:       0,
		OccurredAt:    timestamppb.New(now),
		CorrelationId: requestId,
		CausationId:   &causationId,
		Payload:       payload,
	}); err != nil {
		u.logger.Error("UserUseCase.SentOTP", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AuthUserRegisterResponse{
		Status:  "success",
		Message: "AuthUserRegister",
	}, nil
}
