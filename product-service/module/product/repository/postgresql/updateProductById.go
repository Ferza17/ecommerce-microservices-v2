package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) UpdateProductById(ctx context.Context, product *orm.Product, tx *gorm.DB) (*orm.Product, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.UpdateProductById")
	defer span.End()
	//TODO implement me
	panic("implement me")
}
