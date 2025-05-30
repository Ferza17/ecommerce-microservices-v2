package enum

import (
	"errors"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/notification/v1"
)

type NotificationType string

const (
	NotificationTypeEmailUserOtp NotificationType = "NOTIFICATION.EMAIL.USER.OTP"
)

func (t NotificationType) String() string {
	switch t {
	case
		NotificationTypeEmailUserOtp:
		return string(t)
	default:
		return "unknown"
	}
}

func NotificationTypeParseIntToNotificationType(i int) (NotificationType, error) {
	switch i {
	case int(notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP):
		return NotificationTypeEmailUserOtp, nil

	default:
		return "unknown", errors.New("unknown email type")
	}
}
