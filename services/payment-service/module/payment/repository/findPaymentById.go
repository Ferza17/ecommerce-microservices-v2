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

	if tx == nil {
		tx = r.postgresSQLInfrastructure.GormDB
	}

	if err := tx.WithContext(ctx).
		Preload("PaymentProvider").
		Preload("PaymentItems").
		Where("discarded_at IS NULL").
		First(&payment, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.Error(err.Error())
			return nil, nil
		}
		r.logger.Error(err.Error())
		return nil, err
	}

	return &payment, nil
}
