package orm

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"time"
)

type Provider struct {
	ID          string              `gorm:"column:id;primaryKey;type:uuid;not null"`
	Name        string              `gorm:"column:name;type:varchar(255);not null"`
	Method      enum.ProviderMethod `gorm:"column:method;type:varchar(50);not null"` // Assuming ProviderMethod is mapped to a string
	CreatedAt   time.Time           `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt   time.Time           `gorm:"column:updated_at;type:timestamp"`
	DiscardedAt *time.Time          `gorm:"column:discarded_at;type:timestamp"` // Nullable field
}

// TableName overrides the default table name for GORM
func (Provider) TableName() string {
	return "provider"
}
