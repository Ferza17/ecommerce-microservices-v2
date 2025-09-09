package enum

type ContentType string

const (
	XProtobuf ContentType = "application/x-protobuf"
	JSON      ContentType = "application/json"
)

func (e ContentType) String() string {
	switch e {
	case XProtobuf, JSON:
		return string(e)
	default:
		return "unknown"
	}
}
