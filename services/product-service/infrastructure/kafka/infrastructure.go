package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IKafkaInfrastructure interface {
		PublishWithJsonSchema(ctx context.Context, topic string, key string, value interface{}) error

		Close() error
	}

	kafkaInfrastructure struct {
		producer                *kafka.Producer
		schemaRegistry          schemaregistry.Client
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

func (c *kafkaInfrastructure) Close() error {
	c.producer.Close()
	if err := c.schemaRegistry.Close(); err != nil {
		c.logger.Error(fmt.Sprintf("Failed to close a schema registry client: %v", err))
		return err
	}
	c.logger.Info("Kafka Infrastructure Closed")
	return nil
}

var Set = wire.NewSet(NewKafkaInfrastructure)

func NewKafkaInfrastructure(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IKafkaInfrastructure {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Get().BrokerKafka.Broker1,
		"client.id":         config.Get().UserServiceServiceName,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a kafka producer: %v", err))
		return nil
	}

	schemaRegistry, err := schemaregistry.NewClient(schemaregistry.NewConfig(config.Get().BrokerKafka.SchemaRegistry))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create a schema registry client: %v", err))
		return nil
	}

	return &kafkaInfrastructure{
		producer:                producer,
		schemaRegistry:          schemaRegistry,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
