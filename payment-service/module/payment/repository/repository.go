package repository

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type (
	IPaymentRepository interface {
		FindPaymentById(ctx context.Context, requestId string, id string) (*orm.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, userId string, status enum.PaymentStatus) (*orm.Payment, error)

		UpdatePaymentStatusByIdWithTransaction(ctx context.Context, requestId string, id string, status enum.PaymentStatus, tx *gorm.DB) error
		LockPaymentByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Payment, error)

		// OpenTransactionWithContext
		OpenTransactionWithContext(ctx context.Context) *gorm.DB
	}

	paymentRepository struct {
		postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

// Set is a Wire provider set for Payment repository dependencies
var Set = wire.NewSet(
	NewPaymentRepository,
)

func NewPaymentRepository(
	postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentRepository {
	return &paymentRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}

func (r *paymentRepository) OpenTransactionWithContext(ctx context.Context) *gorm.DB {
	return r.postgresSQLInfrastructure.GormDB().WithContext(ctx).Begin()
}
