package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IUserConsumer interface {
		SnapshotUsersUserCreated(ctx context.Context, message *kafka.Message) error
		SnapshotUsersUserUpdated(ctx context.Context, message *kafka.Message) error
	}

	userConsumer struct {
		kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure
		logger              logger.IZapLogger
		userUseCase         userUseCase.IUserUseCase
	}
)

var Set = wire.NewSet(NewUserConsumer)

func NewUserConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	logger logger.IZapLogger,
	userUseCase userUseCase.IUserUseCase,
) IUserConsumer {
	return &userConsumer{
		kafkaInfrastructure: kafkaInfrastructure,
		logger:              logger,
		userUseCase:         userUseCase,
	}
}
