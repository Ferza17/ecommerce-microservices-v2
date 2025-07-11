package repository

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	paymentRpc "github.com/ferza17/ecommerce-microservices-v2/payment-service/model/rpc/gen/v1/payment"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type (
	IPaymentProviderRepository interface {
		FindPaymentProviderById(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Provider, error)
		FindPaymentProviders(ctx context.Context, requestId string, request *paymentRpc.FindPaymentProvidersRequest, tx *gorm.DB) ([]*orm.Provider, error)
	}

	paymentProviderRepository struct {
		postgresSQLInfrastructure *postgresql.PostgresSQL
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

// Set is a Wire provider set for PaymentProvider repository dependencies
var Set = wire.NewSet(
	NewPaymentProviderRepository,
)

// NewPaymentProviderRepository creates and returns a new instance of IPaymentProviderRepository.
func NewPaymentProviderRepository(
	postgresSQLInfrastructure *postgresql.PostgresSQL,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentProviderRepository {
	return &paymentProviderRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
