package enum

type Event string

const (
	PRODUCT_CREATED Event = "product.created"
	PRODUCT_UPDATED Event = "product.updated"
	PRODUCT_DELETED Event = "product.deleted"
	EVENT_CREATED   Event = "event.created"
)

func (t Event) String() string {
	switch t {
	case PRODUCT_CREATED,
		PRODUCT_UPDATED,
		PRODUCT_DELETED,
		EVENT_CREATED:
		return string(t)
	default:
		return "unknown"
	}
}
