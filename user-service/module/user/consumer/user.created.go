package consumer

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

func (c *userConsumer) UserCreated(ctx context.Context) error {
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
		config.Get().QueueUserCreated,
		config.Get().QueueUserCreated,
		config.Get().ExchangeUser,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := amqpChannel.Consume(
		config.Get().QueueUserCreated,
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
			//request   userRpc.CreateUserRequest
			//requestId string
			)
			carrier := propagation.MapCarrier{}
			for key, value := range d.Headers {
				if key == pkgContext.CtxKeyRequestID {
					//requestId = value.(string)
				}

				if key == pkgContext.CtxKeyAuthorization {
					//ctx = token.SetTokenToContext(ctx, value.(string))
				}

				if strVal, ok := value.(string); ok {
					carrier[key] = strVal
				}
			}
			ctx = otel.GetTextMapPropagator().Extract(context.Background(), carrier)
			//ctx, span := c.telemetryInfrastructure.Tracer(ctx, "AuthConsumer.UserLogin")

			switch d.ContentType {
			case enum.XProtobuf.String():
				//if err = proto.Unmarshal(d.Body, &request); err != nil {
				//	c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
				//	span.End()
				//	continue messages
				//}
			case enum.JSON.String():
				//if err = json.Unmarshal(d.Body, &request); err != nil {
				//	c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
				//	span.End()
				//	continue messages
				//}
			default:
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				//span.End()
				continue messages
			}

			//if _, err = c.userUseCase.CreateUser(ctx, requestId, &request); err != nil {
			//	c.logger.Error(fmt.Sprintf("failed to create user : %v", zap.Error(err)))
			//	span.End()
			//	continue messages
			//}
			//span.End()
		}
	}(msgs)

	<-ctx.Done()
	return nil
}
