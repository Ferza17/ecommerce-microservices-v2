package enum

type service string

const (
	ProductService    service = "product-service"
	UserService       service = "user-service"
	ApiGatewayService service = "api-gateway-service"
)

func (e service) String() string {
	switch e {
	case ProductService,
		ApiGatewayService,
		UserService:
		return string(e)
	default:
		return "unknown"
	}
}
