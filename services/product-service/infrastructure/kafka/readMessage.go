package kafka

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (c *kafkaInfrastructure) ReadMessage(duration time.Duration) (*kafka.Message, error) {
	msg, err := c.consumer.ReadMessage(duration)
	if err != nil {
		if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
			return nil, nil
		}

		c.logger.Error(fmt.Sprintf("failed to read message: %v", err))
		return nil, err
	}
	return msg, nil
}
