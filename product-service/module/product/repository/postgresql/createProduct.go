package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) CreateProduct(ctx context.Context, product *pb.ProductORM, tx *gorm.DB) (string, error) {

	if err := tx.WithContext(ctx).
		Table("products").
		Create(product).
		Error; err != nil {
		return "", err
	}

	return product.Id, nil
}
