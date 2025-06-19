package enum

import (
	"errors"
	notificationRpc "github.com/ferza17/ecommerce-microservices-v2/notification-service/model/rpc/gen/notification/v1"
)

type NotificationType string

const (
	NotificationTypeEmailUserOtp             NotificationType = "NOTIFICATION.EMAIL.USER.OTP"
	NotificationTypeEmailPaymentOrderCreated NotificationType = "NOTIFICATION.EMAIL.PAYMENT.ORDER.CREATED"
)

func (t NotificationType) String() string {
	switch t {
	case
		NotificationTypeEmailUserOtp,
		NotificationTypeEmailPaymentOrderCreated:
		return string(t)
	default:
		return "unknown"
	}
}

func NotificationTypeParseIntToNotificationType(i int) (NotificationType, error) {
	switch i {
	case int(notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_USER_OTP):
		return NotificationTypeEmailUserOtp, nil
	case int(notificationRpc.NotificationTypeEnum_NOTIFICATION_EMAIL_PAYMENT_ORDER_CREATED):
		return NotificationTypeEmailPaymentOrderCreated, nil
	default:
		return "unknown", errors.New("unknown email type")
	}
}
