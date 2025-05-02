package enum

type exchange string

const (
	ProductExchange exchange = "product.exchange"
)

func (e exchange) String() string {
	switch e {
	case ProductExchange:
		return string(e)
	default:
		return "unknown"
	}
}
