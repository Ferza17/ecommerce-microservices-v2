package usecase

import (
	"context"
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/notification"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

func (u *authUseCase) AuthUserRegister(ctx context.Context, requestId string, req *pb.AuthUserRegisterRequest) (*emptypb.Empty, error) {
	var (
		tx  = u.postgresSQL.GormDB.Begin()
		now = time.Now()
	)

	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "AuthUseCase.AuthUserRegister")
	defer span.End()

	// Validate is email already exists
	existedUser, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, req.Email, tx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if existedUser != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("user with this email already exists")))
		return nil, status.Error(codes.AlreadyExists, "User with this email already exists")
	}

	// Validate Role
	role, err := u.rolePostgresqlRepository.FindRoleByName(ctx, requestId, req.Role.String(), tx)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("role not found")))
			return nil, status.Error(codes.NotFound, "role not found")
		}

		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("error hashing password")))
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

	result, err := u.userPostgresqlRepository.CreateUser(ctx, requestId, user, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// Not Intentionally returning password as a response
	result.Password = ""

	otp := util.GenerateOTP()
	if err = u.authRedisRepository.SetOtp(ctx, requestId, otp, result.ID); err != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	reqUserEmailOtp := &notificationRpc.SendOtpEmailNotificationRequest{
		Email:            req.Email,
		Otp:              otp,
		NotificationType: notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP,
	}

	message, err := proto.Marshal(reqUserEmailOtp)
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeNotification, config.Get().QueueNotificationEmailOtpCreated, message); err != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	tx.Commit()

	tx.Commit()
	return nil, nil
}
