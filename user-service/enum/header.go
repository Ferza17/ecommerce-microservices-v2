package enum

type HttpHeader string

const (
	XRequestIDHeader    HttpHeader = "X-Request-Id"
	AuthorizationHeader HttpHeader = "Authorization"
)

func (t HttpHeader) String() string {
	switch t {
	case
		XRequestIDHeader:
		return string(t)
	default:
		return "unknown"
	}
}
