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

	EVENT_CREATED Queue = "event.created"
)

func (t Queue) String() string {
	switch t {
	case PRODUCT_CREATED,
		PRODUCT_UPDATED,
		PRODUCT_DELETED,
		USER_CREATED,
		USER_UPDATED,
		USER_LOGIN,
		CART_CREATED,
		CART_UPDATED,
		EVENT_CREATED:
		return string(t)
	default:
		return "unknown"
	}
}
