package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	pkgMetric "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/metric"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"sync"
)

func (c *notificationEmailConsumer) NotificationEmailOTP(ctx context.Context) error {
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
		config.Get().QueueNotificationEmailOtpCreated,
		config.Get().QueueNotificationEmailOtpCreated,
		config.Get().ExchangeNotification,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	deliveries, err := amqpChannel.Consume(
		config.Get().QueueNotificationEmailOtpCreated,
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

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(deliveries <-chan amqp091.Delivery) {
		defer wg.Done()
	messages:
		for d := range deliveries {
			var (
				request           pb.SendOtpEmailNotificationRequest
				requestId         string
				newCtx, cancelCtx = context.WithTimeout(ctx, 10)
			)
			for key, value := range d.Headers {
				if key == pkgContext.CtxKeyRequestID {
					requestId = value.(string)
					newCtx = pkgContext.SetRequestIDToContext(ctx, requestId)
				}

				if key == pkgContext.CtxKeyAuthorization {
					newCtx = pkgContext.SetTokenAuthorizationToContext(ctx, value.(string))
				}
			}

			newCtx, span := c.telemetryInfrastructure.StartSpanFromRabbitMQHeader(newCtx, d.Headers, "AuthConsumer.UserLogin")
			span.SetAttributes(attribute.String("messaging.destination", config.Get().QueueNotificationEmailOtpCreated))
			span.SetAttributes(attribute.String(pkgContext.CtxKeyRequestID, requestId))

			switch d.ContentType {
			case enum.XProtobuf.String():
				if err = proto.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("requsetID : %s , failed to unmarshal request : %v", requestId, zap.Error(err)))
					pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueNotificationEmailOtpCreated, "failed").Inc()
					span.RecordError(err)
					span.End()
					d.Nack(false, true)
					cancelCtx()
					continue messages
				}
			case enum.JSON.String():
				if err = json.Unmarshal(d.Body, &request); err != nil {
					c.logger.Error(fmt.Sprintf("failed to unmarshal request : %v", zap.Error(err)))
					pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueNotificationEmailOtpCreated, "failed").Inc()
					span.RecordError(err)
					span.End()
					d.Nack(false, true)
					cancelCtx()
					continue messages
				}
			default:
				c.logger.Error(fmt.Sprintf("failed to get request id"))
				pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueNotificationEmailOtpCreated, "failed").Inc()
				span.RecordError(err)
				span.End()
				d.Nack(false, true)
				cancelCtx()
				continue messages
			}

			c.logger.Info(fmt.Sprintf("received a %s message: %s", d.RoutingKey, d.Body))
			if err = c.notificationUseCase.SendNotificationEmailOTP(newCtx, requestId, &request); err != nil {
				pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueNotificationEmailOtpCreated, "failed").Inc()
				span.RecordError(err)
				span.End()
				d.Nack(false, true)
				cancelCtx()
				continue messages
			}

			pkgMetric.RabbitmqMessagesConsumed.WithLabelValues(config.Get().QueueNotificationEmailOtpCreated, "success").Inc()
			span.End()
			cancelCtx()
		}
	}(deliveries)

	wg.Wait()
	return nil
}
