package mongodb

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	INotificationEmailRepository interface {
		FindNotificationTemplateByNotificationType(ctx context.Context, notificationType enum.NotificationType) (*bson.NotificationTemplate, error)
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
	return &notificationEmailRepository{
		mongoDB:                 mongodb,
		logger:                  logger,
		telemetryInfrastructure: telemetryInfrastructure,
	}
}
