package usecase

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *userUseCase) UpdateUserById(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest) (*userRpc.UpdateUserByIdResponse, error) {
	var (
		err error
	)
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "UserUseCase.UpdateUserById")
	defer func() {
		if err != nil {
			span.RecordError(err)
		}
		span.End()
	}()

	now, err := util.GetNowWithTimeZone(pkgContext.CtxValueAsiaJakarta)
	if err != nil {
		u.logger.Error("UserUseCase.UpdateUserById", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id, nil)
	if err != nil {
		return nil, err
	}

	// Partial Update
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			u.logger.Error(fmt.Sprintf("requestId : %s , error hashing password: %v", requestId, err))
			return nil, err
		}
		newPassword := string(hashedPassword)
		req.Password = &newPassword
	}

	if req.Email != nil && *req.Email != user.Email {
		user.Email = *req.Email
	}

	if req.Name != nil && *req.Name != user.Name {
		user.Name = *req.Name
	}

	if req.IsVerified != nil && *req.IsVerified != user.IsVerified {
		user.IsVerified = *req.IsVerified
	}

	payload, err := proto.Marshal(user.ToProto())
	if err != nil {
		u.logger.Error(fmt.Sprintf("requestId : %s , error marshaling payload: %v", requestId, err))
		return nil, err
	}

	// SENT TO EVENT STORE
	if err = u.eventUseCase.AppendEvent(ctx, &pbEvent.Event{
		XId:           primitive.NewObjectID().Hex(),
		AggregateId:   user.ID,
		AggregateType: "users", // TODO: Move To Enum
		EventType:     config.Get().BrokerKafkaTopicUsers.UserUserUpdated,
		Version:       1,
		Timestamp:     timestamppb.New(now),
		SagaId:        requestId,
		Payload:       payload,
	}); err != nil {
		u.logger.Error("UserUseCase.AuthUserRegister", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	user.UpdatedAt = &now
	if err = u.kafkaInfrastructure.PublishWithSchema(ctx, config.Get().BrokerKafkaTopicConnectorSinkPgUser.Users, user.ID, kafka.JSON_SCHEMA, user); err != nil {
		u.logger.Error("UpdateUserById.UpdateUserById", zap.String("requestId", requestId), zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &userRpc.UpdateUserByIdResponse{
		Message: "UpdateUserById",
		Status:  "success",
		Data: &userRpc.UpdateUserByIdResponse_UpdateUserByIdResponseData{
			Id: user.ID,
		},
	}, nil
}
