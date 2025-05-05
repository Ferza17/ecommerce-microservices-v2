package enum

type Event string

const (
	EVENT_CREATED Event = "event.created"
)

func (t Event) String() string {
	switch t {
	case EVENT_CREATED:
		return string(t)
	default:
		return "unknown"
	}
}
