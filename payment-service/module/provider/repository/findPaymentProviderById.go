package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *paymentProviderRepository) FindPaymentProviderById(ctx context.Context, requestId string, id string) (*orm.Provider, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindPaymentProviderById")
	defer span.End()
	var provider *orm.Provider

	// Attempt to find the provider by its ID
	if err := r.postgresSQLInfrastructure.
		GormDB().
		WithContext(ctx).
		First(&provider, "id = ? AND discarded_at IS NULL", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle when no record is found
			r.logger.Error(fmt.Sprintf("Provider not found | ID: %s", id))
			return nil, err
		}
		// Log other unexpected errors
		r.logger.Error(fmt.Sprintf("Failed to find Provider | ID: %s | Error: %v", id, err))
		return nil, err
	}

	return provider, nil
}
