package repository

import (
	"context"
	"errors"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *PaymentRepository) FindPaymentById(ctx context.Context, requestId string, id string) (*orm.Payment, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindPaymentById")
	defer span.End()
	var payment orm.Payment
	// Execute the query with preloading of all foreign keys
	result := r.postgresSQLInfrastructure.GormDB().WithContext(ctx).
		Preload("Provider").
		First(&payment, "id = ?", id)

	// Check and handle errors
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &payment, nil
}
