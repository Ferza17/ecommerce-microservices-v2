package repository

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type (
	IPaymentRepository interface {
		FindPaymentById(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, userId string, status string, tx *gorm.DB) (*orm.Payment, error)

		CreatePayment(ctx context.Context, requestId string, request *orm.Payment, tx *gorm.DB) (string, error)
		CreatePaymentItem(ctx context.Context, paymentItem *orm.PaymentItem, tx *gorm.DB) (string, error)
		UpdatePaymentStatusByIdWithTransaction(ctx context.Context, requestId string, id string, status string, tx *gorm.DB) error
		LockPaymentByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.Payment, error)
	}

	paymentRepository struct {
		postgresSQLInfrastructure *postgresql.PostgresSQL
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

// Set is a Wire provider set for Payment repository dependencies
var Set = wire.NewSet(
	NewPaymentRepository,
)

func NewPaymentRepository(
	postgresSQLInfrastructure *postgresql.PostgresSQL,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentRepository {
	return &paymentRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
