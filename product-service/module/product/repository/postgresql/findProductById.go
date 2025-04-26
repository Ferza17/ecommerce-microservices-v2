package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"gorm.io/gorm"
	"log"
)

func (r *ProductPostgresqlRepository) FindProductById(ctx context.Context, id string, tx *gorm.DB) (*pb.ProductORM, error) {
	product := new(pb.ProductORM)

	if err := tx.WithContext(ctx).
		Table("products").
		Where("id = ?", id).
		First(product).Error; err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return product, nil
}
