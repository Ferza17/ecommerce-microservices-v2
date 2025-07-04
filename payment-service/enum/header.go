package enum

type HttpHeader string

const (
	XDelayHeader HttpHeader = "x-delay"
	XDelayedType HttpHeader = "x-delayed-type"
)

func (t HttpHeader) String() string {
	switch t {
	case
		XDelayHeader,
		XDelayedType:
		return string(t)
	default:
		return "unknown"
	}
}
