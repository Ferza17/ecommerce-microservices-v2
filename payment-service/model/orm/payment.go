package orm

import (
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"time"
)

// Payment represents the GORM model for the Payment table.
type Payment struct {
	ID          string             `gorm:"primaryKey;type:varchar"`                                              // Primary Key
	Code        string             `gorm:"not null"`                                                             // Unique code for the payment
	TotalPrice  float64            `gorm:"not null"`                                                             // Total price of the payment
	Status      enum.PaymentStatus `gorm:"type:paymentstatus;not null"`                                          // Enum field for payment status
	ProviderID  string             `gorm:"type:varchar;not null"`                                                // Foreign key referencing Provider table
	Provider    Provider           `gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Association with Provider
	UserID      string             `gorm:"not null"`                                                             // ID of the user making the payment
	CreatedAt   time.Time          `gorm:"not null;autoCreateTime"`                                              // Default CURRENT_TIMESTAMP
	UpdatedAt   time.Time          `gorm:"not null;autoUpdateTime"`                                              // Auto-update timestamp
	DiscardedAt *time.Time         `gorm:"index"`
}
