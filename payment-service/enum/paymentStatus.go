package enum

import (
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "PENDING"
	PaymentStatusPartial PaymentStatus = "PARTIAL"
	PaymentStatusSuccess PaymentStatus = "SUCCESS"
	PaymentStatusFailed  PaymentStatus = "FAILED"
)

// ProtoToPaymentStatus ProtoToString converts a protobuf PaymentStatus enum to its string representation.
func ProtoToPaymentStatus(status paymentRpc.PaymentStatus) (PaymentStatus, error) {
	switch status {
	case paymentRpc.PaymentStatus_PENDING:
		return PaymentStatusPending, nil
	case paymentRpc.PaymentStatus_PARTIAL:
		return PaymentStatusPartial, nil
	case paymentRpc.PaymentStatus_SUCCESS:
		return PaymentStatusSuccess, nil
	case paymentRpc.PaymentStatus_FAILED:
		return PaymentStatusFailed, nil
	default:
		return "", fmt.Errorf("invalid PaymentStatus: %d", status)
	}
}

// PaymentStatusToProto converts a string representation to its corresponding protobuf PaymentStatus enum.
func PaymentStatusToProto(status PaymentStatus) (paymentRpc.PaymentStatus, error) {
	switch status {
	case PaymentStatusPending:
		return paymentRpc.PaymentStatus_PENDING, nil
	case PaymentStatusPartial:
		return paymentRpc.PaymentStatus_PARTIAL, nil
	case PaymentStatusSuccess:
		return paymentRpc.PaymentStatus_SUCCESS, nil
	case PaymentStatusFailed:
		return paymentRpc.PaymentStatus_FAILED, nil
	default:
		return paymentRpc.PaymentStatus(0), fmt.Errorf("invalid PaymentStatus: %s", status)
	}
}
