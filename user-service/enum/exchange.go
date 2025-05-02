package enum

type exchange string

const (
	USER_EXCHANGE exchange = "user.exchange"
)

func (e exchange) String() string {
	switch e {
	case USER_EXCHANGE:
		return string(e)
	default:
		return "unknown"
	}
}
