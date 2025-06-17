package enum

type HttpHeader string

const (
	XRequestIDHeader HttpHeader = "X-Request-Id"
	XDelayHeader     HttpHeader = "x-delay"
	XDelayedType     HttpHeader = "x-delayed-type"
)

func (t HttpHeader) String() string {
	switch t {
	case
		XDelayHeader,
		XDelayedType,
		XRequestIDHeader:
		return string(t)
	default:
		return "unknown"
	}
}
