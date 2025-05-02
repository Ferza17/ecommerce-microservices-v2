package enum

type event string

const (
	PRODUCT_CREATED event = "product.created"
	PRODUCT_UPDATED event = "product.updated"
	PRODUCT_DELETED event = "product.deleted"
)

func (t event) String() string {
	switch t {
	case PRODUCT_CREATED, PRODUCT_UPDATED, PRODUCT_DELETED:
		return string(t)
	default:
		return "unknown"
	}
}
