package mongodb

import (
	"context"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	eventModel "github.com/ferza17/ecommerce-microservices-v2/user-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventMongoRepository interface {
		FindEventBySagaIDAndAggregateType(ctx context.Context, sagaID string, aggregateType string) (*eventModel.Event, error)
	}

	eventMongoRepository struct {
		mongoDB                 mongodbInfrastructure.IMongoDBInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventMongoDBRepository)

func NewEventMongoDBRepository(
	mongoDB mongodbInfrastructure.IMongoDBInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IEventMongoRepository {
	return &eventMongoRepository{
		mongoDB:                 mongoDB,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
