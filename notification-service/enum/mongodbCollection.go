package enum

type Collection string

const (
	CollectionNotificationTemplate = "notification_templates"
)

func (t Collection) String() string {
	switch t {
	case CollectionNotificationTemplate:
		return string(t)
	default:
		return "unknown"
	}
}
