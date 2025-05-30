package enum

type Queue string

const (
	PRODUCT_CREATED Queue = "product.created"
	PRODUCT_UPDATED Queue = "product.updated"
	PRODUCT_DELETED Queue = "product.deleted"

	CART_CREATED Queue = "cart.created"
	CART_UPDATED Queue = "cart.updated"

	USER_CREATED Queue = "user.created"
	USER_UPDATED Queue = "user.updated"
	USER_LOGIN   Queue = "user.login"
	USER_LOGOUT  Queue = "user.logout"

	NOTIFICATION_EMAIL_OTP Queue = "notification.email.otp"

	EVENT_CREATED Queue = "event.created"
)

func (t Queue) String() string {
	switch t {
	case PRODUCT_CREATED,
		PRODUCT_UPDATED,
		PRODUCT_DELETED,
		USER_CREATED,
		USER_UPDATED,
		CART_CREATED,
		USER_LOGIN,
		USER_LOGOUT,
		CART_UPDATED,
		NOTIFICATION_EMAIL_OTP,
		EVENT_CREATED:
		return string(t)
	default:
		return "unknown"
	}
}
