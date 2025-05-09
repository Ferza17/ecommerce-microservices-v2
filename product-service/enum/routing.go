package enum

type Routing string

const (
	PRODUCT_CREATED Routing = "product.created"
	PRODUCT_UPDATED Routing = "product.updated"
	PRODUCT_DELETED Routing = "product.deleted"
	EVENT_CREATED   Routing = "event.created"
)

func (t Routing) String() string {
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
