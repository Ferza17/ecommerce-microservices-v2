package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

//TODO
// 1. Handle Snapshot
// 2. Handle Confirm
// 3. Handle Compensate

func (c *userConsumer) SnapshotUsersUserUpdated(ctx context.Context, message *kafka.Message) error {
	//TODO implement me
	panic("implement me")
}

func (c *userConsumer) ConfirmSnapshotUsersUserUpdated(ctx context.Context, message *kafka.Message) error {
	//TODO implement me
	panic("implement me")
}

func (c *userConsumer) CompensateSnapshotUsersUserUpdated(ctx context.Context, message *kafka.Message) error {
	//TODO implement me
	panic("implement me")
}
