package enum

type HttpHeader string

const (
	XRequestIDHeader HttpHeader = "X-Request-Id"
	XDelayHeader     HttpHeader = "x-delay"
)

func (t HttpHeader) String() string {
	switch t {
	case
		XDelayHeader,
		XRequestIDHeader:
		return string(t)
	default:
		return "unknown"
	}
}
