package usecase

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"google.golang.org/protobuf/proto"
)

func (u *eventUseCase) AppendEvent(ctx context.Context, request *pb.AppendRequest) error {
	ctx, span := u.telemetryInfrastructure.StartSpanFromContext(ctx, "EventUseCase.AppendEvent")
	defer span.End()

	var (
		requestId = pkgContext.GetRequestIDFromContext(ctx)
	)

	message, err := proto.Marshal(request)
	if err != nil {
		u.logger.Error(fmt.Sprintf("Failed to publish AppendEvent request, requestId: %s, error: %v", requestId, err))
		return err
	}

	if err = u.rabbitmqInfrastructure.PublishFanout(
		ctx,
		requestId,
		config.Get().EventStoreServiceRabbitMQ.ExchangeEventFanout,
		[]string{
			config.Get().EventStoreServiceRabbitMQ.QueueEventEventCreated,
			config.Get().EventStoreServiceRabbitMQ.QueueEventApiGatewayEventCreated,
		},
		message,
	); err != nil {
		u.logger.Error(fmt.Sprintf("Failed to publish AppendEvent request, requestId: %s, error: %v", requestId, err))
	}

	return nil
}
