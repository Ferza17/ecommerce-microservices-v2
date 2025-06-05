package repository

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/payment/v1"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentProviderRepository interface {
		FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentByIdRequest) ([]*orm.Provider, error)
	}

	paymentProviderRepository struct {
		postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

// NewPaymentProviderRepository creates and returns a new instance of IPaymentProviderRepository.
func NewPaymentProviderRepository(
	postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentProviderRepository {
	return &paymentProviderRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
