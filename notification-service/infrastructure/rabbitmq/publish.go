package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error {

	if err := c.channel.ExchangeDeclare(
		exchange,
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

	if err := c.channel.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	carrier := c.telemetryInfrastructure.InjectSpanToTextMapPropagator(ctx)
	headers := amqp091.Table{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: pkgContext.GetTokenAuthorizationFromContext(ctx),
	}
	for k, v := range carrier {
		headers[k] = v
	}

	// Publish message
	if _, err := c.channel.PublishWithDeferredConfirmWithContext(
		ctx,
		exchange,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType:   enum.XProtobuf.String(),
			DeliveryMode:  amqp091.Transient,
			CorrelationId: pkgContext.GetTokenAuthorizationFromContext(ctx),
			Timestamp:     time.Now(),
			Body:          message,
			Headers:       headers,
		},
	); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to publish a message: %v", err))
		return err
	}

	return nil
}
