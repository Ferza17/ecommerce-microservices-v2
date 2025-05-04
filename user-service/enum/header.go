package enum

type headers string

const (
	XRequestId headers = "X-Request-Id"
)

func (h headers) String() string {
	switch h {
	case XRequestId:
		return string(h)
	default:
		return "unknown"
	}
}
