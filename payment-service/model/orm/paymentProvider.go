package orm

import "github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"

// Provider represents the table structure for the provider.
type Provider struct {
	ID     string              `gorm:"primaryKey;type:varchar"`      // id corresponds to VARCHAR PRIMARY KEY
	Name   string              `gorm:"not null"`                     // name corresponds to non-null VARCHAR
	Method enum.ProviderMethod `gorm:"type:providermethod;not null"` // method field is linked to ENUM ProviderMethod
}
