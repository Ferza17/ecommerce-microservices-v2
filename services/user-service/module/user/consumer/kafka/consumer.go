package kafka

import (
	"context"

	kafkaInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/kafka"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pbEvent "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IUserConsumer interface {
		SnapshotUsersUserCreated(ctx context.Context, message *pbEvent.EventEnvelope) error
		SnapshotUsersUserUpdated(ctx context.Context, message *pbEvent.EventEnvelope) error
	}

	userConsumer struct {
		kafkaInfrastructure     kafkaInfrastructure.IKafkaInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
		userUseCase             userUseCase.IUserUseCase
	}
)

var Set = wire.NewSet(NewUserConsumer)

func NewUserConsumer(
	kafkaInfrastructure kafkaInfrastructure.IKafkaInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
	userUseCase userUseCase.IUserUseCase,
) IUserConsumer {
	return &userConsumer{
		kafkaInfrastructure:     kafkaInfrastructure,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
		userUseCase:             userUseCase,
	}
}
