package enum

type Exchange string

const (
	UserExchange         Exchange = "user.exchange"
	EventExchange        Exchange = "event.exchange"
	NotificationExchange Exchange = "notification.exchange"
)

func (e Exchange) String() string {
	switch e {
	case UserExchange,
		NotificationExchange,
		EventExchange:
		return string(e)
	default:
		return "unknown"
	}
}
