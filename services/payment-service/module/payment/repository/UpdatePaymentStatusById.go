package repository

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *paymentRepository) UpdatePaymentStatusByIdWithTransaction(ctx context.Context, requestId string, id string, status string, tx *gorm.DB) error {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPostgresRepository.UpdatePaymentStatusByIdWithTransaction")
	defer span.End()

	if tx == nil {
		tx = r.postgresSQLInfrastructure.GormDB
	}

	// Use the provided transaction to update the status
	result := tx.WithContext(ctx).
		Model(&orm.Payment{}).   // Reference the ORM model
		Where("id = ?", id).     // Match provider by ID
		Update("status", status) // Update the "status" column with the new value

	if result.Error != nil {
		// Log error if the update fails
		r.logger.Error(fmt.Sprintf(
			"Failed to update payment provider status (transaction), requestId: %s, id: %s, error: %v",
			requestId, id, result.Error,
		))
		return result.Error
	}

	// Log successful operation
	r.logger.Info(fmt.Sprintf(
		"Successfully updated payment provider status (transaction), requestId: %s, id: %s, newStatus: %s",
		requestId, id, status,
	))
	return nil
}
