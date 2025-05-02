package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) UpdateProductById(ctx context.Context, product *orm.Product, tx *gorm.DB) (*orm.Product, error) {
	//TODO implement me
	panic("implement me")
}
