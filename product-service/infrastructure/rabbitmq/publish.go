package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
	"time"
)

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error {

	amqpChannel, err := c.amqpConn.Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	defer func(amqpChannel *amqp091.Channel) {
		if err = amqpChannel.Close(); err != nil {
			c.logger.Error(fmt.Sprintf("Failed to close a channel: %v", err))
		}
	}(amqpChannel)

	if err = amqpChannel.ExchangeDeclare(
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

	if err = amqpChannel.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	headers := amqp091.Table{}
	for k, v := range carrier {
		headers[k] = v
	}
	headers[enum.XRequestIDHeader.String()] = requestId

	// Publish message
	if _, err = amqpChannel.PublishWithDeferredConfirmWithContext(
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
