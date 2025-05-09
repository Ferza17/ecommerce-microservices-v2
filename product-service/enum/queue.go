package enum

type Queue string

const (
	QueueProduct  Queue = "product.queue"
	QueueUser     Queue = "user.queue"
	QueueEvent    Queue = "event.queue"
	QueueCommerce Queue = "commerce.queue"
)

func (e Queue) String() string {
	switch e {
	case QueueProduct,
		QueueUser,
		QueueCommerce,
		QueueEvent:
		return string(e)
	default:
		return "unknown"
	}
}
