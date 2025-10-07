package repository

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *paymentProviderRepository) FindPaymentProviderById(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Provider, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "ProviderRepository.FindPaymentProviderById")
	defer span.End()

	if tx == nil {
		tx = r.postgresSQLInfrastructure.GormDB
	}

	var provider orm.Provider
	// Attempt to find the provider by its ID
	if err := tx.
		WithContext(ctx).
		First(&provider, "id = ? AND discarded_at IS NULL", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Handle when no record is found
			r.logger.Error(fmt.Sprintf("Provider not found | ID: %s", id))
			return nil, err
		}
		// Log other unexpected errors
		r.logger.Error(fmt.Sprintf("Failed to find Provider | ID: %s | Error: %v", id, err))
		return nil, err
	}

	return &provider, nil
}
