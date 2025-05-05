package enum

type header string

const (
	XRequestID header = "X-Request-Id"
)

func (e header) String() string {
	switch e {
	case XRequestID:
		return string(e)
	default:
		return "unknown"
	}
}
