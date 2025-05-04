package enum

type Exchange string

const (
	ProductExchange Exchange = "product.exchange"
	UserExchange    Exchange = "user.exchange"
)

func (e Exchange) String() string {
	switch e {
	case ProductExchange,
		UserExchange:
		return string(e)
	default:
		return "unknown"
	}
}
