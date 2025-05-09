package consumer

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/rpc/pb"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *eventConsumer) EventCreated(ctx context.Context) error {

	if err := c.amqpChannel.ExchangeDeclare(
		enum.EventExchange.String(),
		amqp091.ExchangeDirect,
		true,
		false,
		false,
		true,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to declare exchange : %v", zap.Error(err)))
		return err
	}

	if err := c.amqpChannel.QueueBind(
		enum.EVENT_CREATED.String(),
		enum.EVENT_CREATED.String(),
		enum.EventExchange.String(),
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := c.amqpChannel.Consume(
		enum.EVENT_CREATED.String(),
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
	messages:
		for d := range msgs {
			var (
				request   pb.EventStore
				requestId string
				ok        bool
			)

			if requestId, ok = d.Headers[enum.XRequestID.String()].(string); !ok {
				c.logger.Error("failed to get request id")
				continue messages
			}

			if err = proto.Unmarshal(d.Body, &request); err != nil {
				c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
				continue messages
			}

			if _, err = c.eventUseCase.CreateEventStore(ctx, requestId, &request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
				continue messages
			}
		}
	}()

	<-ctx.Done()

	return nil

}
