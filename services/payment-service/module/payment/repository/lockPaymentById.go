package repository

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *paymentRepository) LockPaymentByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Payment, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPostgresRepository.LockPaymentByIdWithTransaction")
	defer span.End()

	var payment orm.Payment

	if tx == nil {
		tx = r.postgresSQLInfrastructure.GormDB
	}

	// Lock the row using SELECT ... FOR UPDATE
	if err := tx.WithContext(ctx).
		Model(&orm.Payment{}).
		Where("id = ?", id).
		Clauses(clause.Locking{
			Strength: clause.LockingStrengthUpdate,
			Options:  clause.LockingOptionsNoWait,
		}).              // Apply row-level lock
		First(&payment). // Select the row
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Log warning if no record is found
			r.logger.Error(fmt.Sprintf(
				"Payment record not found for locking, requestId: %s, id: %s",
				requestId, id,
			))
			return nil, gorm.ErrRecordNotFound
		}

		// Log unexpected errors
		r.logger.Error(fmt.Sprintf(
			"Failed to lock payment record, requestId: %s, id: %s, error: %v",
			requestId, id, err,
		))
		return nil, err
	}

	// Log success of row locking
	r.logger.Info(fmt.Sprintf(
		"Successfully locked payment record, requestId: %s, id: %s",
		requestId, id,
	))
	return &payment, nil
}
