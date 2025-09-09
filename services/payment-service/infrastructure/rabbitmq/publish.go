package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "RabbitMQInfrastructure.Publish")
	defer span.End()

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

	headers := amqp091.Table{
		pkgContext.CtxKeyRequestID:     requestId,
		pkgContext.CtxKeyAuthorization: pkgContext.GetTokenAuthorizationFromContext(ctx),
	}
	for k, v := range c.telemetryInfrastructure.InjectSpanToTextMapPropagator(ctx) {
		headers[k] = v
	}

	// Publish message
	if err := c.channel.PublishWithContext(
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
			MessageId:     fmt.Sprintf("%s:%s", queue, requestId),
			AppId:         config.Get().PaymentServiceServiceName,
		},
	); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to publish a message: %v", err))
		return err
	}

	return nil
}
