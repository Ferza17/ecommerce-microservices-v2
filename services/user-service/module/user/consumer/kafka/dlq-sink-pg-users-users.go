package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *userConsumer) DlqSinkPgUsersUsers(ctx context.Context, message *kafka.Message) error {
	c.logger.Info("error")
	return nil
}
