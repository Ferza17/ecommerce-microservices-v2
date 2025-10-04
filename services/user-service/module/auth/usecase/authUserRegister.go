package usecase

import (
	"context"
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func (u *authUseCase) AuthUserRegister(ctx context.Context, requestId string, req *pb.AuthUserRegisterRequest) (*pb.AuthUserRegisterResponse, error) {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "AuthUseCase.AuthUserRegister")
	defer span.End()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserVerifyOtp", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// Validate is email already exists
	existedUser, err := u.userPostgresqlRepository.FindUserByEmail(ctx, requestId, req.Email, nil)
	if err != nil && err != gorm.ErrRecordNotFound {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if existedUser != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("user with this email already exists")))
		return nil, status.Error(codes.AlreadyExists, "User with this email already exists")
	}

	// Validate Role
	role, err := u.rolePostgresqlRepository.FindRoleByName(ctx, requestId, req.Role.String(), nil)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(errors.New("role not found")))
			return nil, status.Error(codes.NotFound, "role not found")
		}
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
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

	//if err = u.kafkaInfrastructure.PublishWithJsonSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgUser.Users, user.ID, user); err != nil {
	//	u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
	//	return nil, status.Error(codes.Internal, "internal server error")
	//}

	payload, err := proto.Marshal(user.ToProto())
	if err != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// SENT TO EVENT STORE
	if err = u.eventUseCase.AppendEvent(ctx, &pbEvent.Event{
		XId:           primitive.NewObjectID().Hex(),
		AggregateId:   user.ID,
		AggregateType: "users", // TODO: Move To Enum
		EventType:     config.Get().BrokerKafkaTopicUsers.UserUserCreated,
		Version:       1,
		Timestamp:     timestamppb.New(now),
		SagaId:        requestId,
		//Payload:       &pbEvent.Event_User{User: user.ToProto()},
		Payload: payload,
	}); err != nil {
		u.logger.Error("AuthUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	// SENT TO NOTIFICATION

	if err = u.SentOTP(ctx, requestId, user.ToProto()); err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.AuthUserRegisterResponse{
		Status:  "success",
		Message: "AuthUserRegister",
	}, nil
}
