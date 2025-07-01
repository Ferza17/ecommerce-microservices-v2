package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
	"log"
	"time"
)

func (r *ProductPostgresqlRepository) FindProductById(ctx context.Context, id string, tx *gorm.DB) (*orm.Product, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := r.telemetryInfrastructure.Tracer(ctxTimeout, "Repository.FindProductById")
	defer span.End()
	product := new(orm.Product)
	if err := tx.WithContext(ctxTimeout).
		Table("products").
		Where("id = ?", id).
		First(product).Error; err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return product, nil
}
