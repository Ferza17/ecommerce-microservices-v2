package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *authConsumer) SnapshotUsersUserLogin(ctx context.Context, message *kafka.Message) error {
	//TODO: Handle for topic SnapshotUsersUserLogin
	c.logger.Info("SnapshotUsersUserLogin")
	c.logger.Info(fmt.Sprintf("message: %v", message))
	return nil
}
