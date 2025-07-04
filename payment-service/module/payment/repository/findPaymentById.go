package repository

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *paymentRepository) FindPaymentById(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Payment, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "PaymentPostgresRepository.FindPaymentById")
	defer span.End()
	var payment orm.Payment
	// Execute the query with preloading of all foreign keys
	result := tx.WithContext(ctx).
		Preload("Provider").
		First(&payment, "id = ?", id)

	// Check and handle errors

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &payment, nil
}
