package enum

type service string

const (
	ProductService service = "product-service"
	UserService    service = "user-service"
)

func (e service) String() string {
	switch e {
	case ProductService,
		UserService:
		return string(e)
	default:
		return "unknown"
	}
}
