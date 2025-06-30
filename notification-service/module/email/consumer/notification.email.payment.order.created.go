package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *notificationEmailConsumer) NotificationEmailPaymentOrderCreated(ctx context.Context) error {
	amqpChannel, err := c.rabbitmqInfrastructure.GetConnection().Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		config.Get().ExchangeNotification,
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
		config.Get().QueueNotificationEmailPaymentOrderCreated,
		config.Get().QueueNotificationEmailPaymentOrderCreated,
		config.Get().ExchangeNotification,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	msgs, err := amqpChannel.Consume(
		config.Get().QueueNotificationEmailPaymentOrderCreated,
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
				request   notificationRpc.SendEmailPaymentOrderCreateRequest
				requestId string
			)
			carrier := propagation.MapCarrier{}
			for key, value := range d.Headers {
				if key == pkgContext.CtxKeyRequestID {
					requestId = value.(string)
				}

				if strVal, ok := value.(string); ok {
					carrier[key] = strVal
				}
			}
			ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
			ctx, span := c.telemetryInfrastructure.Tracer(ctx, "Consumer.NotificationLoginCreated")

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

			if err = c.notificationUseCase.SendNotificationEmailPaymentOrderCreated(ctx, requestId, &request); err != nil {
				span.End()
				c.logger.Error(fmt.Sprintf("failed to login user : %v", zap.Error(err)))
				continue messages
			}
			span.End()
		}

	}(msgs)
	<-ctx.Done()

	if err = amqpChannel.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a channel: %v", err))
		return err
	}
	return nil
}
