package enum

type headers string

const (
	XRequestId headers = "X-REQUEST-ID"
)

func (h headers) String() string {
	switch h {
	case XRequestId:
		return string(h)
	default:
		return "unknown"
	}
}
