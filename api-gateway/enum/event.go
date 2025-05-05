package enum

type Event string

const (
	PRODUCT_CREATED Event = "product.created"
	PRODUCT_UPDATED Event = "product.updated"
	PRODUCT_DELETED Event = "product.deleted"

	USER_CREATED Event = "user.created"
	USER_UPDATED Event = "user.updated"

	EVENT_CREATED Event = "event.created"
)

func (t Event) String() string {
	switch t {
	case PRODUCT_CREATED,
		PRODUCT_UPDATED,
		PRODUCT_DELETED,
		USER_CREATED,
		USER_UPDATED,
		EVENT_CREATED:
		return string(t)
	default:
		return "unknown"
	}
}
