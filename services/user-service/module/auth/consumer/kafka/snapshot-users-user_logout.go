package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *authConsumer) SnapshotUsersUserLogout(ctx context.Context, message *kafka.Message) error {
	c.logger.Info("SnapshotUsersUserLogout")
	c.logger.Info(fmt.Sprintf("message: %v", message))
	return nil
}

func (c *authConsumer) ConfirmSnapshotUsersUserLogout(ctx context.Context, message *kafka.Message) error {
	c.logger.Info("SnapshotUsersUserLogout")
	c.logger.Info(fmt.Sprintf("message: %v", message))
	return nil
}

func (c *authConsumer) CompensateSnapshotUsersUserLogout(ctx context.Context, message *kafka.Message) error {
	c.logger.Info("SnapshotUsersUserLogout")
	c.logger.Info(fmt.Sprintf("message: %v", message))
	return nil
}
