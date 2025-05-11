package enum

import (
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/pb"
)

type NotificationType string

const (
	NotificationTypeEmailUserLogin   NotificationType = "NOTIFICATION.EMAIL.USER.LOGIN"
	NotificationTypeEmailUserCreated NotificationType = "NOTIFICATION.EMAIL.USER.CREATED"
)

func (t NotificationType) String() string {
	switch t {
	case
		NotificationTypeEmailUserLogin,
		NotificationTypeEmailUserCreated:
		return string(t)
	default:
		return "unknown"
	}
}

func NotificationTypeParseIntToNotificationType(i int) (NotificationType, error) {
	switch i {
	case int(pb.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_CREATED):
		return NotificationTypeEmailUserCreated, nil
	case int(pb.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_LOGIN):
		return NotificationTypeEmailUserLogin, nil
	default:
		return "unknown", errors.New("unknown notification type")
	}
}
