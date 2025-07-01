package postgres

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/model/orm"
	"gorm.io/gorm"
	"time"
)

func (r *ProductPostgresqlRepository) UpdateProductById(ctx context.Context, product *orm.Product, tx *gorm.DB) (*orm.Product, error) {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := r.telemetryInfrastructure.Tracer(ctxTimeout, "Repository.UpdateProductById")
	defer span.End()
	//TODO implement me
	panic("implement me")
}
