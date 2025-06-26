package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/token"

	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *userConsumer) UserUpdated(ctx context.Context) error {
	amqpChannel, err := c.rabbitmqInfrastructure.GetConnection().Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		config.Get().ExchangeUser,
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
		config.Get().QueueUserUpdated,
		config.Get().QueueUserUpdated,
		config.Get().ExchangeUser,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	deliveries, err := amqpChannel.ConsumeWithContext(
		ctx,
		config.Get().QueueUserUpdated,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to consume : %v", zap.Error(err)))
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func(deliveries <-chan amqp091.Delivery) {
		defer cancel()
	messages:
		for d := range deliveries {
			var (
				request   userRpc.UpdateUserByIdRequest
				requestId string
			)
			carrier := propagation.MapCarrier{}
			for key, value := range d.Headers {
				if key == enum.XRequestIDHeader.String() {
					requestId = value.(string)
				}

				if key == enum.AuthorizationHeader.String() {
					ctx = token.SetTokenToContext(ctx, value.(string))
				}

				if strVal, ok := value.(string); ok {
					carrier[key] = strVal
				}
			}
			ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
			ctx, span := c.telemetryInfrastructure.Tracer(ctx, "AuthConsumer.UserLogin")

			switch d.ContentType {
			case enum.XProtobuf.String():
				if err = proto.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
					span.End()
					continue messages
				}
			case enum.JSON.String():
				if err = json.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
					span.End()
					continue messages
				}
			default:
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				span.End()
				continue messages
			}

			c.logger.Info(fmt.Sprintf("received a %s message: %s", d.RoutingKey, d.Body))
			if _, err = c.userUseCase.UpdateUserById(ctx, requestId, &request); err != nil {
				c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
				span.End()
				continue messages
			}
			span.End()
		}
	}(deliveries)

	<-ctx.Done()
	return nil
}
