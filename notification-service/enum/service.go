package enum

type service string

const (
	ProductService      service = "product-service"
	UserService         service = "user-service"
	NotificationService service = "notification-service"
)

func (e service) String() string {
	switch e {
	case ProductService,
		NotificationService,
		UserService:
		return string(e)
	default:
		return "unknown"
	}
}
