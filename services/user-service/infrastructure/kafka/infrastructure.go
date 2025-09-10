package kafka

import (
	"context"
	"github.com/IBM/sarama"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IKafkaInfrastructure interface {
		Publish(ctx context.Context, message *sarama.ProducerMessage) error
		Consume(topic string) ([]sarama.PartitionConsumer, error)
	}

	KafkaInfrastructure struct {
		config                  *sarama.Config
		logger                  logger.IZapLogger
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
	}
)

var Set = wire.NewSet(NewKafkaInfrastructure)

func NewKafkaInfrastructure(
	logger logger.IZapLogger,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
) IKafkaInfrastructure {
	config := sarama.NewConfig()

	// Consumer Config
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	// Producer Config
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	config.Version = sarama.MaxVersion

	return &KafkaInfrastructure{
		config:                  config,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
