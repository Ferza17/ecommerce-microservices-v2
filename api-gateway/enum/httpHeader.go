package enum

type HttpHeader string

const (
	AcceptHeader        HttpHeader = "Accept"
	AuthorizationHeader HttpHeader = "Authorization"
	ContentTypeHeader   HttpHeader = "Content-Type"
	XCSRFTokenHeader    HttpHeader = "X-CSRF-Token"
	XRequestIDHeader    HttpHeader = "X-Request-Id"
)

func (t HttpHeader) String() string {
	switch t {
	case AcceptHeader,
		AuthorizationHeader,
		ContentTypeHeader,
		XCSRFTokenHeader,
		XRequestIDHeader:
		return string(t)
	default:
		return "unknown"
	}
}
