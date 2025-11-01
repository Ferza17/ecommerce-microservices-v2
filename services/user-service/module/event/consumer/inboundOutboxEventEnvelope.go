package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *eventConsumer) InboundOutboxEventEnvelope(ctx context.Context, message *kafka.Message) error {
	c.logger.Error(fmt.Sprintf("InboundOutboxEventEnvelope : %s", string(message.Value)))
	return nil
}
