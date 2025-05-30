package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/event/v1"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/user/v1"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *userUseCase) UpdateUserById(ctx context.Context, requestId string, req *userRpc.UpdateUserByIdRequest) (*userRpc.UpdateUserByIdResponse, error) {
	var (
		err        error
		tx         = u.userPostgresqlRepository.OpenTransactionWithContext(ctx)
		eventStore = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       enum.ProductService.String(),
			EventType:     enum.USER_UPDATED.String(),
			Status:        enum.SUCCESS.String(),
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
			eventStore.Status = enum.FAILED.String()
		}

		if err = u.rabbitmqInfrastructure.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

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

	result, err := u.userPostgresqlRepository.UpdateUserByIdWithTransaction(ctx, requestId, req, tx)
	if err != nil {
		tx.Rollback()
		u.logger.Error(fmt.Sprintf("requestId : %s , error updating user: %v", requestId, err))
		return nil, err
	}

	tx.Commit()
	return &userRpc.UpdateUserByIdResponse{
		Id: result,
	}, nil
}
