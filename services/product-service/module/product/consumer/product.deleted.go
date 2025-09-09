package consumer

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

func (c *productConsumer) ProductDeleted(ctx context.Context, d *amqp091.Delivery) error {
	return nil
}
