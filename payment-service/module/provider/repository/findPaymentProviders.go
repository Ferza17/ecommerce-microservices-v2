package repository

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
)

func (r *paymentProviderRepository) FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest) ([]*orm.Provider, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindPaymentProviders")
	defer span.End()
	var providers []*orm.Provider

	// Query the database for all providers
	result := r.postgresSQLInfrastructure.GormDB().WithContext(ctx).Find(&providers)
	if result.Error != nil {
		// Log the error and requestId for traceability
		r.logger.Error(fmt.Sprintf("Failed to fetch payment providers, requestId : %s , error : %v", requestId, result.Error))
		return nil, result.Error
	}

	return providers, nil
}
