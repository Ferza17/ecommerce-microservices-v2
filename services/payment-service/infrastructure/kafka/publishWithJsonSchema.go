package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"

	pkgContext "github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/context"
)

func (c *kafkaInfrastructure) PublishWithJsonSchema(ctx context.Context, topic string, key string, value interface{}) error {
	var (
		headers = []kafka.Header{
			{
				Key:   pkgContext.CtxKeyRequestID,
				Value: []byte(pkgContext.GetRequestIDFromContext(ctx)),
			},
		}
		deliveryChan = make(chan kafka.Event, 1)
	)
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "KafkaInfrastructure.PublishWithJsonSchema")
	defer span.End()

	if pkgContext.GetTokenAuthorizationFromContext(ctx) != "" {
		headers = append(headers, kafka.Header{
			Key:   pkgContext.CtxKeyAuthorization,
			Value: []byte(pkgContext.GetTokenAuthorizationFromContext(ctx)),
		})
	}

	// Add tracing headers
	carrier := c.telemetryInfrastructure.InjectSpanToTextMapPropagator(ctx)
	for k, v := range carrier {
		headers = append(headers, kafka.Header{
			Key:   k,
			Value: []byte(v),
		})
	}

	// Create JSON Schema serializer with proper configuration

	// Ensure the value is properly structured
	payload, err := c.jsonSerializer.Serialize(topic, value)
	if err != nil {
		c.logger.Error(fmt.Sprintf("failed to serialize message: %v", err))
		return err
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:     []byte(key),
		Value:   payload,
		Headers: headers,
	}

	if err = c.producer.Produce(message, deliveryChan); err != nil {
		c.logger.Error(fmt.Sprintf("failed to publish message to topic %s: %v", topic, err))
		return err
	}

	// Wait for delivery confirmation
	select {
	case e := <-deliveryChan:
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			c.logger.Error(fmt.Sprintf("delivery failed: %v", m.TopicPartition.Error))
			return m.TopicPartition.Error
		} else {
			c.logger.Info(fmt.Sprintf("delivered message to topic %s [%d] at offset %v",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))
		}
	case <-time.After(10 * time.Second):
		return fmt.Errorf("delivery timeout for topic %s", topic)
	}

	return nil

}
