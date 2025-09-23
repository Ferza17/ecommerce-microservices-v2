package worker

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaTaskQueue struct {
	Ctx     context.Context
	Message *kafka.Message
	Handler func(ctx context.Context, message *kafka.Message) error
}
