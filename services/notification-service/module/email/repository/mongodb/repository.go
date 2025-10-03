package mongodb

import (
	"context"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	pb "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/v1/notification"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	INotificationEmailRepository interface {
		FindNotificationTemplateByNotificationType(ctx context.Context, requestId string, notificationType pb.NotificationTypeEnum) (*bson.NotificationTemplate, error)
	}

	notificationEmailRepository struct {
		mongoDB                 mongodbInfrastructure.IMongoDBInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewNotificationEmailRepository)

func NewNotificationEmailRepository(
	mongodb mongodbInfrastructure.IMongoDBInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	logger logger.IZapLogger) INotificationEmailRepository {
	c := &notificationEmailRepository{
		mongoDB:                 mongodb,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
	return c
}
