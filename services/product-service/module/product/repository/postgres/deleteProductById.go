package postgres

import (
	"context"
	"gorm.io/gorm"
	"time"
)

func (r *ProductPostgresqlRepository) DeleteProductById(ctx context.Context, id string, tx *gorm.DB) error {
	var (
		ctxTimeout, cancel = context.WithTimeout(ctx, 5*time.Second)
	)
	defer cancel()
	ctxTimeout, span := r.telemetryInfrastructure.StartSpanFromContext(ctxTimeout, "Repository.DeleteProductById")
	defer span.End()
	//TODO implement me
	panic("implement me")
}
