package rabbitmq

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

//func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange enum.Exchange, event enum.Event, message []byte) error {
//
//	amqpChannel, err := c.amqpConn.Channel()
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
//		return err
//	}
//
//	defer func(amqpChannel *amqp091.Channel) {
//		err = amqpChannel.Close()
//		if err != nil {
//			c.logger.Error(fmt.Sprintf("Failed to close a channel: %v", err))
//		}
//	}(amqpChannel)
//
//	if err = amqpChannel.ExchangeDeclare(
//		exchange.String(),
//		amqp091.ExchangeTopic,
//		true,
//		false,
//		false,
//		false,
//		nil,
//	); err != nil {
//		c.logger.Error(fmt.Sprintf("Failed to declare an exchange: %s", err.Error()))
//		return err
//	}
//
//	if err = amqpChannel.QueueBind(event.String(), "", exchange.String(), false, nil); err != nil {
//		c.logger.Error(fmt.Sprintf("Failed to bind a queue: %s", err.Error()))
//	}
//
//	if err = amqpChannel.Confirm(false); err != nil {
//		c.logger.Error(fmt.Sprintf("Failed to confirm a message: %s", err.Error()))
//		return err
//	}
//
//	// Publish message
//	if err = amqpChannel.PublishWithContext(ctx,
//		exchange.String(),
//		event.String(),
//		false,
//		false,
//		amqp091.Publishing{
//			ContentType:  "application/json",
//			DeliveryMode: amqp091.Transient,
//			Timestamp:    time.Now(),
//			Body:         message,
//			Headers: map[string]interface{}{
//				enum.XRequestIDHeader.String(): requestId,
//			},
//		},
//	); err != nil {
//		c.logger.Error(fmt.Sprintf("Failed to publish a message: %v", err))
//		return err
//	}
//
//	return nil
//}

func (c *RabbitMQInfrastructure) Publish(ctx context.Context, requestId string, exchange enum.Exchange, event enum.Event, message []byte) error {

	amqpChannel, err := c.amqpConn.Channel()
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to create a channel: %v", err))
		return err
	}

	if err = amqpChannel.ExchangeDeclare(
		enum.ProductExchange.String(),
		amqp091.ExchangeTopic, // type
		true,                  // durable
		false,                 // auto-delete
		false,
		true,
		nil,
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to declare exchange : %v", zap.Error(err)))
		return err
	}

	defer func(amqpChannel *amqp091.Channel) {
		if err = amqpChannel.Close(); err != nil {
			c.logger.Error(fmt.Sprintf("Failed to close a channel: %v", err))
		}
	}(amqpChannel)

	q, err := amqpChannel.QueueDeclare(
		event.String(),
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		c.logger.Error(fmt.Sprintf("Failed to declare a queue: %v", err))
		return err
	}

	// Publish message
	if _, err = amqpChannel.PublishWithDeferredConfirmWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "text/plain",
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
