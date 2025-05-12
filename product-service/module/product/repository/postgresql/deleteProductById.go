package postgresql

import (
	"context"
	"gorm.io/gorm"
)

func (r *ProductPostgresqlRepository) DeleteProductById(ctx context.Context, id string, tx *gorm.DB) error {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.DeleteProductById")
	defer span.End()
	//TODO implement me
	panic("implement me")
}
