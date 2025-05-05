package enum

type Exchange string

const (
	ProductExchange Exchange = "product.exchange"
	EventExchange   Exchange = "event.exchange"
)

func (e Exchange) String() string {
	switch e {
	case ProductExchange,
		EventExchange:
		return string(e)
	default:
		return "unknown"
	}
}
