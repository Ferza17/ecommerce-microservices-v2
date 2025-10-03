package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	pkgContext "github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/context"
)

func (c *kafkaInfrastructure) PublishWithSchema(ctx context.Context, topic string, key string, schemaType SchemaType, value interface{}) error {
	ctx, span := c.telemetryInfrastructure.StartSpanFromContext(ctx, "KafkaInfrastructure.PublishWithSchema")
	defer span.End()

	var (
		headers []kafka.Header
		payload []byte
		err     error
	)

	switch schemaType {
	case JSON_SCHEMA:
		headers, payload, err = c.jsonSerializer.SerializeWithHeaders(topic, value)
		if err != nil {
			c.logger.Error(fmt.Sprintf("failed to serialize message: %v", err))
			return err
		}
		break
	case PROTOBUF_SCHEMA:
		headers, payload, err = c.protobufSerializer.SerializeWithHeaders(topic, value)
		if err != nil {
			c.logger.Error(fmt.Sprintf("failed to serialize message: %v", err))
			return err
		}
		break
	case AVRO:
		headers, payload, err = c.avroSerializer.SerializeWithHeaders(topic, value)
		if err != nil {
			c.logger.Error(fmt.Sprintf("failed to serialize message: %v", err))
			return err
		}
		break
	default:
		c.logger.Error("SCHEMA TYPE NOT SUPPORTED")
		return errors.New("SCHEMA TYPE NOT SUPPORTED")
	}

	headers = append(headers, kafka.Header{
		Key:   pkgContext.CtxKeyRequestID,
		Value: []byte(pkgContext.GetRequestIDFromContext(ctx)),
	})

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
	
	if err = c.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:     []byte(key),
		Value:   payload,
		Headers: headers,
	}, nil); err != nil {
		c.logger.Error(fmt.Sprintf("failed to publish message to topic %s: %v", topic, err))
		return err
	}

	return nil
}
