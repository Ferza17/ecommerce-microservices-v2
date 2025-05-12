package mongodb

import (
	"context"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

type (
	IEventRepository interface {
		CreateEvent(ctx context.Context, requestId string, req *bson.Event) (string, error)
	}

	eventRepository struct {
		mongoDB                 mongodbInfrastructure.IMongoDBInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  pkg.IZapLogger
	}
)

func NewEventRepository(
	mongodb mongodbInfrastructure.IMongoDBInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger pkg.IZapLogger) IEventRepository {
	return &eventRepository{
		mongoDB:                 mongodb,
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
