package enum

type exchange string

const (
	EventExchange exchange = "event.exchange"
)

func (e exchange) String() string {
	switch e {
	case EventExchange:
		return string(e)
	default:
		return "unknown"
	}
}
