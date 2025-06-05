package enum

import (
	"fmt"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
)

type ProviderMethod string

const (
	ProviderMethodBank           ProviderMethod = "BANK"
	ProviderMethodCryptoCurrency ProviderMethod = "CRYPTO_CURRENCY"
	ProviderMethodDebit          ProviderMethod = "DEBIT"
	ProviderMethodCredit         ProviderMethod = "CREDIT"
	ProviderMethodCashOnDelivery ProviderMethod = "CASH_ON_DELIVERY"
)

// ProtoToProviderMethod converts a protobuf ProviderMethod enum to its string representation.
func ProtoToProviderMethod(method paymentRpc.ProviderMethod) ProviderMethod {
	switch method {
	case paymentRpc.ProviderMethod_BANK:
		return ProviderMethodBank
	case paymentRpc.ProviderMethod_CRYPTO_CURRENCY:
		return ProviderMethodCryptoCurrency
	case paymentRpc.ProviderMethod_DEBIT:
		return ProviderMethodDebit
	case paymentRpc.ProviderMethod_CREDIT:
		return ProviderMethodCredit
	case paymentRpc.ProviderMethod_CASH_ON_DELIVERY:
		return ProviderMethodCashOnDelivery
	default:
		return "UNKNOWN"
	}
}

// ProviderMethodToProto converts a string representation of ProviderMethod to its corresponding protobuf enum.
func ProviderMethodToProto(method ProviderMethod) (paymentRpc.ProviderMethod, error) {
	switch method {
	case ProviderMethodBank:
		return paymentRpc.ProviderMethod_BANK, nil
	case ProviderMethodCryptoCurrency:
		return paymentRpc.ProviderMethod_CRYPTO_CURRENCY, nil
	case ProviderMethodDebit:
		return paymentRpc.ProviderMethod_DEBIT, nil
	case ProviderMethodCredit:
		return paymentRpc.ProviderMethod_CREDIT, nil
	case ProviderMethodCashOnDelivery:
		return paymentRpc.ProviderMethod_CASH_ON_DELIVERY, nil
	default:
		return paymentRpc.ProviderMethod(0), fmt.Errorf("invalid ProviderMethod: %s", method)
	}
}
