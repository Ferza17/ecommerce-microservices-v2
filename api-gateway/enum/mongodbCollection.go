package enum

type Collection string

const (
	CollectionUserEvent = "event-store-gateway-event"
)

func (t Collection) String() string {
	switch t {
	case CollectionUserEvent:
		return string(t)
	default:
		return "unknown"
	}
}
