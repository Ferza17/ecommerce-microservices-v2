package kafka

import (
	"context"

	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	authUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IAuthConsumer interface {
		SnapshotUsersUserLogin(ctx context.Context, message *pbEvent.EventEnvelope) error
		SnapshotUsersUserLogout(ctx context.Context, message *pbEvent.EventEnvelope) error
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
