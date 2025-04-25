package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) FindProductById(ctx context.Context, id int64, tx *gorm.DB) (*pb.ProductORM, error) {
	return nil, nil
}
