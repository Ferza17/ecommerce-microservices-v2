package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/avrov2"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/jsonschema"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IKafkaInfrastructure interface {
		PublishWithSchema(ctx context.Context, topic string, key string, schemaType SchemaType, value interface{}) error
		Publish(ctx context.Context, topic string, key string, value []byte) error

		SetupTopics(topics []string) error
		ReadMessage(duration time.Duration) (*kafka.Message, error)
		Close() error
	}

	kafkaInfrastructure struct {
		producer                *kafka.Producer
		consumer                *kafka.Consumer
		jsonSerializer          *jsonschema.Serializer
		protobufSerializer      *protobuf.Serializer
		avroSerializer          *avrov2.Serializer
		registryClient          schemaregistry.Client
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}

	SchemaType string
)

const (
	JSON_SCHEMA     SchemaType = "JSON"
	PROTOBUF_SCHEMA SchemaType = "PROTOBUF"
	AVRO            SchemaType = "AVRO"
)

func (c *kafkaInfrastructure) Close() error {
	c.producer.Close()
	c.consumer.Close()
	c.registryClient.Close()
	c.jsonSerializer.Close()
	c.avroSerializer.Close()
	return nil
}

var Set = wire.NewSet(NewKafkaInfrastructure)

func NewKafkaInfrastructure(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IKafkaInfrastructure {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": config.Get().BrokerKafka.Broker1,
		"client.id":         config.Get().ConfigServiceUser.ServiceName,

		// Consumer config
		"auto.offset.reset":     "earliest",
		"group.id":              config.Get().ConfigServiceUser.ServiceName,
		"session.timeout.ms":    10000,
		"heartbeat.interval.ms": 3000,
	}

	producer, err := kafka.NewProducer(configMap)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a kafka producer: %v", err))
		return nil
	}

	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a kafka consumer: %v", err))
		return nil
	}

	registryClient, err := schemaregistry.NewClient(schemaregistry.NewConfig(config.Get().BrokerKafka.SchemaRegistry))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a schema registry client: %v", err))
		return nil
	}

	jsonSerializer, err := jsonschema.NewSerializer(registryClient, serde.ValueSerde, jsonschema.NewSerializerConfig())
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create kafka serializer: %v", err))
		return nil
	}

	protobufSerializer, err := protobuf.NewSerializer(registryClient, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create kafka deserializer: %v", err))
		return nil
	}

	avroSerializer, err := avrov2.NewSerializer(registryClient, serde.ValueSerde, avrov2.NewSerializerConfig())
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create kafka deserializer: %v", err))
		return nil
	}

	return &kafkaInfrastructure{
		producer:                producer,
		consumer:                consumer,
		jsonSerializer:          jsonSerializer,
		protobufSerializer:      protobufSerializer,
		avroSerializer:          avroSerializer,
		registryClient:          registryClient,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
