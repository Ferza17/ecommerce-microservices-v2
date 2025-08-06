package mongodb

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	telemetryInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/temporal"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
)

type (
	INotificationEmailRepository interface {
		FindNotificationTemplateByNotificationType(ctx context.Context, requestId string, notificationType enum.NotificationType) (*bson.NotificationTemplate, error)
	}

	notificationEmailRepository struct {
		mongoDB                 mongodbInfrastructure.IMongoDBInfrastructure
		telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure
		temporal                temporal.ITemporalInfrastructure
		logger                  logger.IZapLogger
	}
)

var Set = wire.NewSet(NewNotificationEmailRepository)

func NewNotificationEmailRepository(
	mongodb mongodbInfrastructure.IMongoDBInfrastructure,
	telemetryInfrastructure telemetryInfrastructure.ITelemetryInfrastructure,
	temporal temporal.ITemporalInfrastructure,
	logger logger.IZapLogger) INotificationEmailRepository {
	c := &notificationEmailRepository{
		mongoDB:                 mongodb,
		logger:                  logger,
		temporal:                temporal,
		telemetryInfrastructure: telemetryInfrastructure,
	}
	c.temporal = c.temporal.
		RegisterActivity(c.FindNotificationTemplateByNotificationType)
	return c
}
