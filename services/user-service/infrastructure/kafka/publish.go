package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
	"google.golang.org/protobuf/proto"
)

func (c *kafkaInfrastructure) Publish(ctx context.Context, topic string, key string, schemaType SchemaType, value interface{}) error {
	var (
		headers = []kafka.Header{
			{
				Key:   pkgContext.CtxKeyRequestID,
				Value: []byte(pkgContext.GetRequestIDFromContext(ctx)),
			},
		}
	)
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "KafkaInfrastructure.Publish")
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

	var (
		payload []byte
		err     error
	)
	switch schemaType {
	case JSON_SCHEMA:
		payload, err = json.Marshal(value)
		if err != nil {
			return err
		}
	case PROTOBUF_SCHEMA:
		v, ok := value.(proto.Message)
		if !ok {
			c.logger.Error(fmt.Sprintf("failed to marshal value: %v", value))
			return fmt.Errorf("value is not a proto message")
		}
		payload, err = proto.Marshal(v)
		if err != nil {
			return err
		}
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

	if err = c.producer.Produce(message, nil); err != nil {
		c.logger.Error(fmt.Sprintf("failed to publish message to topic %s: %v", topic, err))
		return err
	}
	return nil
}
