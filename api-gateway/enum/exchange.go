package enum

type Exchange string

const (
	ProductExchange Exchange = "product.exchange"
	UserExchange    Exchange = "user.exchange"
	EventExchange   Exchange = "event.exchange"
)

func (e Exchange) String() string {
	switch e {
	case ProductExchange,
		UserExchange,
		EventExchange:
		return string(e)
	default:
		return "unknown"
	}
}
