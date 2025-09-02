package mongodb

import (
	"context"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/enum"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/model/bson"
	pkgLogger "github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg/logger"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IEventRepository interface {
		CreateEvent(ctx context.Context, requestId string, event *bson.Event) (*bson.Event, error)
	}

	eventRepository struct {
		collection              *mongo.Collection
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
		collection:              mongodb.GetCollection(enum.DatabaseEventStore, "events"),
		telemetryInfrastructure: telemetryInfrastructure,
		logger:                  logger,
	}
}
