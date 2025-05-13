package postgresql

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	"gorm.io/gorm"
)

func (r *userPostgresqlRepository) FindUserByIdWithTransaction(ctx context.Context, requestId string, id string, tx *gorm.DB) (*orm.User, error) {
	ctx, span := r.telemetryInfrastructure.Tracer(ctx, "Repository.FindUserByIdWithTransaction")
	defer span.End()
	//TODO implement me
	panic("implement me")
}
