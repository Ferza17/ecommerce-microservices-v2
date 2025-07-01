package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
	"time"
)

func (r *ProductPostgresqlRepository) CreateProduct(ctx context.Context, product *orm.Product, tx *gorm.DB) (string, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := r.telemetryInfrastructure.Tracer(ctxTimeout, "Repository.CreateProduct")
	defer span.End()
	if err := tx.WithContext(ctxTimeout).
		Table("products").
		Create(product).
		Error; err != nil {
		return "", err
	}

	return product.ID, nil
}
