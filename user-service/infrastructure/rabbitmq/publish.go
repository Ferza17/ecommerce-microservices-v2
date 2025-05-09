package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange enum.Exchange, queue enum.Queue, message []byte) error {
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
		exchange.String(),
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
		queue.String(),
		"",
		exchange.String(),
		false,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue : %v", zap.Error(err)))
		return err
	}

	// Publish message
	if _, err = amqpChannel.PublishWithDeferredConfirmWithContext(
		ctx,
		exchange.String(),
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Transient,
			Timestamp:    time.Now(),
			Body:         message,
			Headers: map[string]interface{}{
				enum.XRequestIDHeader.String(): requestId,
			},
		},
	); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to publish a message: %v", err))
		return err
	}

	return nil
}
