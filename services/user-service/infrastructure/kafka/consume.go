package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
)

func (c *KafkaInfrastructure) Consume(topic string) ([]sarama.PartitionConsumer, error) {
	consumer, err := sarama.NewConsumer([]string{
		config.Get().BrokerKafka.Broker1,
	}, c.config)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to create kafka consumer: %v", err))
	}

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to get partitions for topic %s: %v", topic, err))
	}
	c.logger.Info(fmt.Sprintf("partitions for topic %s: %v", topic, partitions))

	consumerPartitions := make([]sarama.PartitionConsumer, len(partitions))
	for _, partition := range partitions {
		consumerPartition, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			c.logger.Error(fmt.Sprintf("failed to consume partition %d of topic %s: %v", partition, topic, err))
			return nil, err
		}
		consumerPartitions = append(consumerPartitions, consumerPartition)
	}

	return consumerPartitions, nil
}
