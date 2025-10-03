package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IAuthConsumer interface {
		SnapshotUsersUserLogin(ctx context.Context, message *kafka.Message) error
		ConfirmSnapshotUsersUserLogin(ctx context.Context, message *kafka.Message) error
		CompensateSnapshotUsersUserLogin(ctx context.Context, message *kafka.Message) error

		SnapshotUsersUserLogout(ctx context.Context, message *kafka.Message) error
		ConfirmSnapshotUsersUserLogout(ctx context.Context, message *kafka.Message) error
		CompensateSnapshotUsersUserLogout(ctx context.Context, message *kafka.Message) error
	}

	authConsumer struct {
		kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure
		logger              logger.IZapLogger
		authUseCase         authUseCase.IAuthUseCase
	}
)

var Set = wire.NewSet(NewAuthConsumer)

func NewAuthConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	logger logger.IZapLogger,
	authUseCase authUseCase.IAuthUseCase) IAuthConsumer {
	return &authConsumer{
		kafkaInfrastructure: kafkaInfrastructure,
		logger:              logger,
		authUseCase:         authUseCase,
	}
}
