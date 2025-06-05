package orm

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"time"
)

// Provider represents the table structure for the provider.
type Provider struct {
	ID          string              `gorm:"primaryKey;type:varchar"`      // id corresponds to VARCHAR PRIMARY KEY
	Name        string              `gorm:"not null"`                     // name corresponds to non-null VARCHAR
	Method      enum.ProviderMethod `gorm:"type:providermethod;not null"` // method field is linked to ENUM ProviderMethod
	CreatedAt   time.Time           `gorm:"not null;autoCreateTime"`      // Default CURRENT_TIMESTAMP
	UpdatedAt   time.Time           `gorm:"not null;autoUpdateTime"`      // Auto-update timestamp
	DiscardedAt *time.Time          `gorm:"index"`
}
