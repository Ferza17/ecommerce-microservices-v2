package enum

type Event string

const (
	PRODUCT_CREATED Event = "product.created"
	PRODUCT_UPDATED Event = "product.updated"
	PRODUCT_DELETED Event = "product.deleted"
)

func (t Event) String() string {
	switch t {
	case PRODUCT_CREATED, PRODUCT_UPDATED, PRODUCT_DELETED:
		return string(t)
	default:
		return "unknown"
	}
}
