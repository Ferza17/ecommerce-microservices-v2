package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *userConsumer) SnapshotUsersUserUpdated(ctx context.Context, message *kafka.Message) error {
	//TODO implement me
	panic("implement me")
}
