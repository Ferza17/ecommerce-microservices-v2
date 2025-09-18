package worker

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQTaskQueue struct {
	QueueName string
	Ctx       context.Context
	Delivery  *amqp091.Delivery
	Handler   func(ctx context.Context, d *amqp091.Delivery) error
}
