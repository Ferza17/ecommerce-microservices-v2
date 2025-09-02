package rabbitmq

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func (c RabbitMQInfrastructure) Consume(ctx context.Context, queue string) (<-chan amqp091.Delivery, error) {
	deliveries, err := c.channel.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack (set to false for manual acknowledgment)
		false, // exclusive
		false, // no-local
		true,  // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("failed to register consumer for queue %s: %v", queue, err)
		return nil, err
	}
	return deliveries, nil
}
