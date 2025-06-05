package repository

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
)

type (
	IPaymentRepository interface {
		FindPaymentById(ctx context.Context, requestId string, id string) (*orm.Payment, error)
		FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, userId string, status enum.PaymentStatus) (*orm.Payment, error)
	}

	PaymentRepository struct {
		postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		logger                    logger.IZapLogger
	}
)

func NewPaymentRepository(
	postgresSQLInfrastructure postgresql.IPostgreSQLInfrastructure,
	telemetryInfrastructure telemetry.ITelemetryInfrastructure,
	logger logger.IZapLogger,
) IPaymentRepository {
	return &PaymentRepository{
		postgresSQLInfrastructure: postgresSQLInfrastructure,
		telemetryInfrastructure:   telemetryInfrastructure,
		logger:                    logger,
	}
}
