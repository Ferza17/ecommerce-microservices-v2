package consumer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *eventConsumer) DlqSinkMongoEventsPaymentEventStores(ctx context.Context, message *kafka.Message) error {
	c.logger.Error(fmt.Sprintf("DLQ Sink Mongo Events User Event Stores : %s", string(message.Value)))
	return nil
}
