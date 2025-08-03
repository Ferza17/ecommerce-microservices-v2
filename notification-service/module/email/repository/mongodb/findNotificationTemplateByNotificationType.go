package mongodb

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/enum"
	model "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/bson"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (r *notificationEmailRepository) FindNotificationTemplateByNotificationType(ctx context.Context, notificationType enum.NotificationType) (*model.NotificationTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Increase to 30 seconds or more
	defer cancel()
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "NotificationMongoDBRepository.FindNotificationTemplateByNotificationType")
	defer span.End()
	resp := new(model.NotificationTemplate)
	filter := bson.M{"type": notificationType.String()}
	if err := r.mongoDB.
		GetCollection(enum.DatabaseNotification, enum.CollectionNotificationTemplate).
		FindOne(ctx, filter).
		Decode(resp); err != nil {
		r.logger.Error(fmt.Sprintf("failed to find email template by email type : %v", err))
		return nil, err
	}
	return resp, nil
}
