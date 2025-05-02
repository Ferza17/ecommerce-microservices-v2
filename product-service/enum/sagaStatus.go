package enum

type SagaStatus string

const (
	PENDING  SagaStatus = "pending"
	SUCCESS  SagaStatus = "success"
	FAILED   SagaStatus = "failed"
	ROLLBACK SagaStatus = "rollback"
)

func (s SagaStatus) String() string {
	switch s {
	case PENDING, SUCCESS, FAILED, ROLLBACK:
		return string(s)
	default:
		return "unknown"
	}
}
