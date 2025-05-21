package enum

type ContextKey string

const (
	ContextKeyUserID      ContextKey = "ctx.key.user.id"
	ContextKeyAccessToken ContextKey = "ctx.key.access.token"
)

func (t ContextKey) String() string {
	switch t {
	case ContextKeyUserID,
		ContextKeyAccessToken:
		return string(t)
	default:
		return "unknown"
	}
}
