package rabbitmq

import (
	"fmt"

	"go.uber.org/zap"
)

// SetupQueue declares the queues and exchanges
func (c *RabbitMQInfrastructure) SetupQueue(exchange string, topic string, queue string) error {
	// Declare exchange (optional - depending on your setup)
	err := c.channel.ExchangeDeclare(
		exchange,
		topic,
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to declare exchange : %v", zap.Error(err)))
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// Declare queues
	if _, err = c.channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	); err != nil {
		c.logger.Error(fmt.Sprintf("failed to declare queue : %v", zap.Error(err)))
		return fmt.Errorf("failed to declare queue %s: %w", queue, err)
	}

	// Bind queue to exchange (optional - depending on your setup)
	err = c.channel.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil,
	)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to bind queue %s : %v", queue, err))
		return fmt.Errorf("failed to bind queue %s: %w", queue, err)
	}

	return nil
}
