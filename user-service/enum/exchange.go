package enum

type Exchange string

const (
	USER_EXCHANGE Exchange = "user.exchange"
	EventExchange Exchange = "event.exchange"
)

func (e Exchange) String() string {
	switch e {
	case USER_EXCHANGE, EventExchange:
		return string(e)
	default:
		return "unknown"
	}
}
