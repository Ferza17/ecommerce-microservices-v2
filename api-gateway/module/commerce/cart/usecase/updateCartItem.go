package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	commerceRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/commerce/v1"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/event/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *CartUseCase) UpdateCartItemById(ctx context.Context, requestId string, req *commerceRpc.UpdateCartItemByIdRequest) (*commerceRpc.UpdateCartItemByIdResponse, error) {
	var (
		err        error = nil
		eventStore       = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().CommerceServiceName,
			EventType:     config.Get().QueueCommerceCartUpdated,
			Status:        config.Get().CommonSagaStatusPending,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
		event = config.Get().QueueCommerceCartUpdated
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "UpdateCartItemById")

	defer func(err error, eventStore *eventRpc.EventStore) {
		defer span.End()
		if err != nil {
			eventStore.Status = config.Get().CommonSagaStatusFailed
		}

		eventStoreMessage, err := proto.Marshal(eventStore)
		if err != nil {
			u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
			return
		}

		if err = u.rabbitMQ.Publish(ctx, requestId, config.Get().ExchangeEvent, config.Get().QueueEventCreated, eventStoreMessage); err != nil {
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

	// Need to add this line for matching with Nestjs Constructor
	message, err := json.Marshal(map[string]interface{}{
		"pattern": event,
		"data":    req,
	})
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return nil, err
	}

	if err = u.rabbitMQ.Publish(ctx, requestId, config.Get().ExchangeCommerce, event, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return nil, err
	}

	return &commerceRpc.UpdateCartItemByIdResponse{}, nil
}
