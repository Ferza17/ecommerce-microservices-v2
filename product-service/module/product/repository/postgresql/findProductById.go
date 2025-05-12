package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
	"log"
)

func (r *ProductPostgresqlRepository) FindProductById(ctx context.Context, id string, tx *gorm.DB) (*orm.Product, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindProductById")
	defer span.End()
	product := new(orm.Product)
	if err := tx.WithContext(ctx).
		Table("products").
		Where("id = ?", id).
		First(product).Error; err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return product, nil
}
