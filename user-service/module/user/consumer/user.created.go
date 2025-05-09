package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *userConsumer) UserCreated(ctx context.Context) error {
	amqpChannel, err := c.rabbitmqInfrastructure.GetConnection().Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		enum.USER_EXCHANGE.String(),
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

	if err = amqpChannel.QueueBind(
		enum.USER_CREATED.String(),
		enum.USER_CREATED.String(),
		enum.USER_EXCHANGE.String(),
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := amqpChannel.Consume(
		enum.USER_CREATED.String(),
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func(deliveries <-chan amqp091.Delivery) {
		defer cancel()
	messages:
		for d := range deliveries {
			var (
				request   pb.CreateUserRequest
				requestId string
				ok        bool
			)
			if requestId, ok = d.Headers[enum.XRequestIDHeader.String()].(string); !ok {
				c.logger.Error("failed to get request id")
				continue messages
			}

			switch d.ContentType {
			case enum.XProtobuf.String():
				if err = proto.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
					continue messages
				}
			case enum.JSON.String():
				if err = json.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
					continue messages
				}
			default:
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				continue messages
			}

			c.logger.Info(fmt.Sprintf("received a %s message: %s", d.RoutingKey, d.Body))

			if _, err = c.userUseCase.CreateUser(ctx, requestId, &request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
				continue messages
			}
		}
	}(msgs)

	<-ctx.Done()
	return nil
}
