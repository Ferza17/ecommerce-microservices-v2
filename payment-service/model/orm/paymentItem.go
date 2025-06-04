package orm

import "time"

// PaymentItem represents the GORM model for the PaymentItem table.
type PaymentItem struct {
	ID        string    `gorm:"primaryKey;type:varchar"` // Primary key for the PaymentItem
	ProductID string    `gorm:"not null"`                // Foreign key to the product
	Amount    float64   `gorm:"not null"`                // Monetary amount
	Qty       int32     `gorm:"not null"`                // Quantity of products
	CreatedAt time.Time `gorm:"not null;autoCreateTime"` // Record creation timestamp
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"` // Record update timestamp
}
