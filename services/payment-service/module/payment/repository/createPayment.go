package repository

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *paymentRepository) CreatePayment(ctx context.Context, requestId string, request *orm.Payment, tx *gorm.DB) (string, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPostgresRepository.CreatePayment")
	defer span.End()

	// Validate that the transaction object exists
	if tx == nil {
		tx = r.postgresSQLInfrastructure.GormDB
	}

	// Create the Payment record
	if err := tx.WithContext(ctx).Create(request).Error; err != nil {
		r.logger.Error(fmt.Sprintf(
			"Failed to create Payment record, requestId: %s, error: %v",
			requestId, err,
		))
		return "", fmt.Errorf("failed to create payment record: %w", err)
	}

	// Return the payment ID
	return request.ID, nil
}
