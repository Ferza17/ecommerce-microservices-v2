package repository

import (
	"context"
	"fmt"

	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"gorm.io/gorm"
)

func (r *paymentProviderRepository) FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest, tx *gorm.DB) ([]*orm.Provider, error) {
	ctx, span := r.telemetryInfrastructure.StartSpanFromContext(ctx, "ProviderRepository.FindPaymentProviders")
	defer span.End()
	var providers []*orm.Provider

	if tx == nil {
		tx = r.postgresSQLInfrastructure.GormDB
	}

	// Query the database for all providers
	result := tx.WithContext(ctx).Find(&providers)
	if result.Error != nil {
		// Log the error and requestId for traceability
		r.logger.Error(fmt.Sprintf("Failed to fetch payment providers, requestId : %s , error : %v", requestId, result.Error))
		return nil, result.Error
	}

	return providers, nil
}
