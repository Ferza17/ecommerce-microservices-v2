package repository

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *paymentRepository) CreatePaymentItem(ctx context.Context, paymentItem *orm.PaymentItem, tx *gorm.DB) (string, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPostgresRepository.CreatePaymentItem")
	defer span.End()

	// Validate that the transaction object exists
	if tx == nil {
		return "", fmt.Errorf("transaction object cannot be nil")
	}

	// Insert the PaymentItem record
	if err := tx.WithContext(ctx).Create(paymentItem).Error; err != nil {
		r.logger.Error(fmt.Sprintf(
			"Failed to create PaymentItem, paymentItemId: %s, PaymentID: %s, error: %v",
			paymentItem.ID, paymentItem.PaymentID, err,
		))
		return "", fmt.Errorf("failed to create payment item: %w", err)
	}

	// Return the PaymentItem ID
	return paymentItem.ID, nil
}
