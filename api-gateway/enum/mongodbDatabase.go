package enum

type Database string

const (
	DatabaseEventStore = "event-store"
)

func (t Database) String() string {
	switch t {
	case DatabaseEventStore:
		return string(t)
	default:
		return "unknown"
	}
}
