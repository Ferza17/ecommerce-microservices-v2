package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *userUseCase) UpdateUserById(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest) (*userRpc.UpdateUserByIdResponse, error) {
	var (
		err        error
		tx         = u.postgresSQLInfrastructure.GormDB.Begin()
		eventStore = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ServiceName,
			EventType:     config.Get().QueueUserUpdated,
			Status:        config.Get().CommonSagaStatusSuccess,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UpdateUserById")

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
			eventStore.Status = config.Get().CommonSagaStatusFailed
		}

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	user, err := u.userPostgresqlRepository.FindUserById(ctx, requestId, req.Id, tx)
	if err != nil {
		return nil, err
	}

	// Partial Update
	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
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

	result, err := u.userPostgresqlRepository.UpdateUserById(ctx, requestId, user, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error updating user: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &userRpc.UpdateUserByIdResponse{
		Id: result.ID,
	}, nil
}
