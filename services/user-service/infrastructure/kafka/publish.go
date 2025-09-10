package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
)

func (c *KafkaInfrastructure) Publish(ctx context.Context, message *sarama.ProducerMessage) error {
	producer, err := sarama.NewSyncProducer([]string{}, c.config)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to create kafka producer: %v", err))
		return err
	}
	defer producer.Close()

	if _, _, err = producer.SendMessage(message); err != nil {
		c.logger.Error(fmt.Sprintf("failed to publish message to topic %s: %v", message.Topic, err))
		return err
	}

	return nil
}
