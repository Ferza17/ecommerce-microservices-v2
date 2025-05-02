package enum

type headers string

const (
	X_REQUEST_ID headers = "X-REQUEST-ID"
)

func (h headers) String() string {
	switch h {
	case X_REQUEST_ID:
		return string(h)
	default:
		return "unknown"
	}
}
