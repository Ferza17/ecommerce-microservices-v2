package enum

type Database string

const (
	DatabaseNotification = "notification"
)

func (t Database) String() string {
	switch t {
	case DatabaseNotification:
		return string(t)
	default:
		return "unknown"
	}
}
