package mongodb

import (
	"context"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	pkgLogger "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
)

type (
	IEventRepository interface {
		CreateEvent(ctx context.Context, requestId string, req *bson.Event) (string, error)
	}

	eventRepository struct {
		mongoDB                 mongodbInfrastructure.IMongoDBInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkgLogger.IZapLogger
	}
)

var Set = wire.NewSet(NewEventRepository)

func NewEventRepository(
	mongodb mongodbInfrastructure.IMongoDBInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkgLogger.IZapLogger) IEventRepository {
	return &eventRepository{
		mongoDB:                 mongodb,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
