package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange string, queue string, message []byte) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "RabbitMQInfrastructure.Publish")

	amqpChannel, err := c.amqpConn.Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	defer func(amqpChannel *amqp091.Channel) {
		span.AddEvent(queue)
		defer span.End()
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

	carrier := c.telemetryInfrastructure.InjectSpanToTextMapPropagator(ctx)
	headers := amqp091.Table{}
	for k, v := range carrier {
		headers[k] = v
	}
	headers[pkgContext.CtxKeyRequestID] = requestId

	// Publish message
	if err = amqpChannel.PublishWithContext(
		ctx,
		exchange,
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType:  enum.XProtobuf.String(),
			DeliveryMode: amqp091.Transient,
			Timestamp:    time.Now(),
			Body:         message,
			Headers:      headers,
		},
	); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to publish a message: %v", err))
		return err
	}

	return nil
}
