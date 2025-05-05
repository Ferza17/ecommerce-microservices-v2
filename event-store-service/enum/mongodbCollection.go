package enum

type Collection string

const (
	CollectionEvent = "event"
)

func (t Collection) String() string {
	switch t {
	case CollectionEvent:
		return string(t)
	default:
		return "unknown"
	}
}
