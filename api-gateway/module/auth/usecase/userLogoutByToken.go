package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/event/v1"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"

	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *authUseCase) UserLogoutByToken(ctx context.Context, requestId string, req *userRpc.UserLogoutByTokenRequest) (*userRpc.UserLogoutByTokenResponse, error) {
	var (
		err        error = nil
		eventStore       = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       enum.UserService.String(),
			EventType:     enum.USER_LOGOUT.String(),
			Status:        enum.PENDING.String(),
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UseCase.UserLogoutByToken")
	defer span.End()

	// Validation If User Exists
	user, err := u.authService.FindUserByToken(ctx, requestId, &userRpc.FindUserByTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("error finding user by email and password: %s", err.Error()))
		return nil, err
	}

	if user == nil {
		u.logger.Error(fmt.Sprintf("user not found"))
		return nil, fmt.Errorf("user not found")
	}

	defer func(err error, eventStore *eventRpc.EventStore) {
		if err != nil {
			eventStore.Status = enum.FAILED.String()
		}

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
			return
		}

		if err = u.rabbitMQ.Publish(ctx, requestId, enum.EventExchange, enum.EVENT_CREATED, eventStoreMessage); err != nil {
			u.logger.Error(fmt.Sprintf("error creating product event store: %s", err.Error()))
			return
		}
	}(err, eventStore)

	payload, err := util.ConvertStructToProtoStruct(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error converting struct to proto struct: %s", err.Error()))
		return nil, err
	}
	eventStore.Payload = payload
	message, err := proto.Marshal(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return nil, err
	}
	if err = u.rabbitMQ.Publish(ctx, requestId, enum.UserExchange, enum.USER_LOGOUT, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return nil, err
	}
	return &userRpc.UserLogoutByTokenResponse{}, nil
}
