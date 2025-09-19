package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/jsonschema"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IKafkaInfrastructure interface {
		PublishWithJsonSchema(ctx context.Context, topic string, key string, value interface{}) error
		Publish(ctx context.Context, topic string, key string, value []byte) error
		SetupTopics(topics []string) error
		ReadMessage(duration time.Duration) (*kafka.Message, error)
		Close() error
	}

	kafkaInfrastructure struct {
		producer                *kafka.Producer
		consumer                *kafka.Consumer
		jsonSerializer          *jsonschema.Serializer
		protobufDeserializer    *protobuf.Deserializer
		registryClient          schemaregistry.Client
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

func (c *kafkaInfrastructure) Close() error {
	c.producer.Close()
	c.consumer.Close()
	c.registryClient.Close()
	c.jsonSerializer.Close()
	return nil
}

var Set = wire.NewSet(NewKafkaInfrastructure)

func NewKafkaInfrastructure(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IKafkaInfrastructure {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": config.Get().BrokerKafka.Broker1,
		"client.id":         config.Get().ConfigServiceNotification.ServiceName,

		// Consumer config
		"auto.offset.reset":     "earliest",
		"group.id":              config.Get().ConfigServiceNotification.ServiceName,
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

	protobufDeserializer, err := protobuf.NewDeserializer(registryClient, serde.ValueSerde, protobuf.NewDeserializerConfig())
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create kafka deserializer: %v", err))
		return nil
	}

	return &kafkaInfrastructure{
		producer:                producer,
		consumer:                consumer,
		jsonSerializer:          jsonSerializer,
		protobufDeserializer:    protobufDeserializer,
		registryClient:          registryClient,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
