package usecase

import (
	"context"
	rabbitmqInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/event"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventUseCase interface {
		AppendEvent(ctx context.Context, request *pb.AppendRequest) error
	}

	eventUseCase struct {
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		rabbitmqInfrastructure  rabbitmqInfrastructure.IRabbitMQInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventUseCase)

func NewEventUseCase(
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	rabbitmqInfrastructure rabbitmqInfrastructure.IRabbitMQInfrastructure,
	logger logger.IZapLogger,
) IEventUseCase {
	return &eventUseCase{
		telemetryInfrastructure: telemetryInfrastructure,
		rabbitmqInfrastructure:  rabbitmqInfrastructure,
		logger:                  logger,
	}
}
