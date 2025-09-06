package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func (c *RabbitMQInfrastructure) PublishFanout(ctx context.Context, requestId string, exchange string, queue string, message []byte) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "RabbitMQInfrastructure.PublishFanout")
	defer span.End()

	if err := c.channel.ExchangeDeclare(
		exchange,
		amqp091.ExchangeFanout,
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
	if err := c.channel.PublishWithContext(
		ctx,
		exchange,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType:   enum.XProtobuf.String(),
			CorrelationId: pkgContext.GetTokenAuthorizationFromContext(ctx),
			DeliveryMode:  amqp091.Transient,
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
