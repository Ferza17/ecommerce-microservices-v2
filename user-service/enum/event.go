package enum

type Event string

const (
	USER_CREATED  Event = "user.created"
	USER_UPDATED  Event = "user.updated"
	EVENT_CREATED Event = "event.created"
)

func (t Event) String() string {
	switch t {
	case USER_CREATED,
		USER_UPDATED,
		EVENT_CREATED:
		return string(t)
	default:
		return "unknown"
	}
}
