package enum

type ElasticsearchIndex string

const (
	ElasticsearchIndexProduct ElasticsearchIndex = "products"
)

func (e ElasticsearchIndex) String() string {
	switch e {
	case ElasticsearchIndexProduct:
		return string(e)
	default:
		return "unknown"
	}
}
