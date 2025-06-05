package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/model/orm"
	"gorm.io/gorm"
)

func (r *PaymentRepository) FindPaymentByUserIdAndStatus(ctx context.Context, requestId string, userId string, status enum.PaymentStatus) (*orm.Payment, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindPaymentById")
	defer span.End()

	// Variable to hold the payment record
	var payment orm.Payment

	// Execute query with preloading
	result := r.postgresSQLInfrastructure.GormDB().WithContext(ctx).
		Preload("Provider"). // Preload the Provider association
		Where("user_id = ? AND status = ?", userId, status).
		First(&payment)

	// Handle query errors
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.logger.Error(fmt.Sprintf("requestId : %s  err Record not found", requestId))
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		r.logger.Error(fmt.Sprintf("requestId : %s  err : %v", requestId, result.Error))
		return nil, result.Error // Return the error if something went wrong
	}

	return &payment, nil
}
