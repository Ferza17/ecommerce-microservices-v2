package enum

type meta string

const (
	XRequestID meta = "X-REQUEST-ID"
)

func (e meta) String() string {
	switch e {
	case XRequestID:
		return string(e)
	default:
		return "unknown"
	}
}
