package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *roleConsumer) DlqSinkPgUsersRoles(ctx context.Context, message *kafka.Message) error {
	//TODO implement me
	panic("implement me")
}
