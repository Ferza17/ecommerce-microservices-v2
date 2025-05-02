package consumer

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"go.uber.org/zap"
	"log"
)

func (c *userConsumer) UserUpdated(ctx context.Context) error {
	q, err := c.amqpChannel.QueueDeclare(
		enum.USER_UPDATED.String(),
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to serve queue ", zap.Error(err)))
		return err
	}

	msgs, err := c.amqpChannel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		for d := range msgs {
			log.Printf("Received a %s message: %s", enum.USER_UPDATED.String(), d.Body)
		}
	}()

	<-ctx.Done()
	return nil
}
