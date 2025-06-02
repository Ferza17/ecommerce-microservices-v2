package usecase

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/config"
	eventRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/event/v1"
	productRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/product/v1"

	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/util"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *ProductUseCase) CreateProduct(ctx context.Context, requestId string, req *productRpc.CreateProductRequest) (*productRpc.CreateProductResponse, error) {
	var (
		err        error = nil
		eventStore       = &eventRpc.EventStore{
			RequestId:     requestId,
			Service:       config.Get().ProductServiceName,
			EventType:     config.Get().QueueProductCreated,
			Status:        config.Get().CommonSagaStatusPending,
			PreviousState: nil,
			CreatedAt:     timestamppb.Now(),
			UpdatedAt:     timestamppb.Now(),
		}
	)
	ctx, span := u.telemetryInfrastructure.Tracer(ctx, "CreateProduct")

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

	message, err := proto.Marshal(req)
	if err != nil {
		u.logger.Error(fmt.Sprintf("error marshaling message: %s", err.Error()))
		return nil, err
	}

	if err = u.rabbitMQ.Publish(ctx, requestId, config.Get().ExchangeProduct, config.Get().QueueProductCreated, message); err != nil {
		u.logger.Error(fmt.Sprintf("error publishing message to rabbitmq: %s", err.Error()))
		return nil, err
	}

	return &productRpc.CreateProductResponse{}, nil
}
