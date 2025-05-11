package mongodb

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	mongodbInfrastructure "github.com/ferza17/ecommerce-microservices-v2/notification-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg"
)

type (
	INotificationRepository interface {
		FindNotificationTemplateByNotificationType(ctx context.Context, notificationType enum.NotificationType) (*bson.NotificationTemplate, error)
	}

	notificationRepository struct {
		mongoDB mongodbInfrastructure.IMongoDBInfrastructure
		logger  pkg.IZapLogger
	}
)

func NewEventRepository(mongodb mongodbInfrastructure.IMongoDBInfrastructure, logger pkg.IZapLogger) INotificationRepository {
	return &notificationRepository{
		mongoDB: mongodb,
		logger:  logger,
	}
}
