package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/pb"
	"gorm.io/gorm"
)

type (
	IProductPostgresqlRepository interface {
		FindProductById(ctx context.Context, id int64, tx *gorm.DB) (*pb.ProductORM, error)
	}

	ProductPostgresqlRepository struct {
	}
)

func NewProductPostgresqlRepository() *ProductPostgresqlRepository {
	return &ProductPostgresqlRepository{}
}
