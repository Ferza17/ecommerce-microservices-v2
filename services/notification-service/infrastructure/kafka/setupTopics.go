package kafka

import "fmt"

func (c *kafkaInfrastructure) SetupTopics(topics []string) error {
	err := c.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to subscribe to topics: %v", err))
		return err
	}
	return nil
}
