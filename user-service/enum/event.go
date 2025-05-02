package enum

type event string

const (
	USER_CREATED event = "user.created"
	USER_UPDATED event = "user.updated"
)

func (t event) String() string {
	switch t {
	case USER_CREATED, USER_UPDATED:
		return string(t)
	default:
		return "unknown"
	}
}
