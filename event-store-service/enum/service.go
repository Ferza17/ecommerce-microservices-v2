package enum

type service string

const (
	ProductService    service = "product-service"
	UserService       service = "user-service"
	ApiGatewayService service = "api-gateway-service"
	EventStoreService service = "event-store-service"
)

func (e service) String() string {
	switch e {
	case ProductService,
		ApiGatewayService,
		EventStoreService,
		UserService:
		return string(e)
	default:
		return "unknown"
	}
}
