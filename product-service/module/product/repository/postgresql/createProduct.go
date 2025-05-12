package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) CreateProduct(ctx context.Context, product *orm.Product, tx *gorm.DB) (string, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.CreateProduct")
	defer span.End()
	if err := tx.WithContext(ctx).
		Table("products").
		Create(product).
		Error; err != nil {
		return "", err
	}

	return product.ID, nil
}
