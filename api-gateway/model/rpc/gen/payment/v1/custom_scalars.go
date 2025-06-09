package gen

import (
	"fmt"
	"io"
	"strconv"
)

// MarshalGQL and UnmarshalGQL for PaymentStatus
func (e PaymentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *PaymentStatus) UnmarshalGQL(v interface{}) error {
	strValue, ok := v.(string)
	if !ok {
		return fmt.Errorf("PaymentStatus must be a string, got %T", v)
	}

	enumValue, exists := PaymentStatus_value[strValue]
	if !exists {
		return fmt.Errorf("invalid PaymentStatus: %s", strValue)
	}

	*e = PaymentStatus(enumValue)
	return nil
}

// MarshalGQL and UnmarshalGQL for ProviderMethod
func (e ProviderMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *ProviderMethod) UnmarshalGQL(v interface{}) error {
	strValue, ok := v.(string)
	if !ok {
		return fmt.Errorf("ProviderMethod must be a string, got %T", v)
	}

	enumValue, exists := ProviderMethod_value[strValue]
	if !exists {
		return fmt.Errorf("invalid ProviderMethod: %s", strValue)
	}

	*e = ProviderMethod(enumValue)
	return nil
}
