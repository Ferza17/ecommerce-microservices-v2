package kafka

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func (c *kafkaInfrastructure) CommitMessage(msg *kafka.Message) error {
	_, err := c.consumer.CommitMessage(msg)
	if err != nil {
		return err
	}
	return nil
}