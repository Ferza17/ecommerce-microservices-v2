package orm

import "time"

// PaymentItem represents the GORM model for the PaymentItem table.
type PaymentItem struct {
	ID          string     `gorm:"primaryKey;type:varchar"`     // Primary key
	PaymentID   string     `gorm:"not null;type:varchar;index"` // Foreign key to Payment
	ProductID   string     `gorm:"not null;type:varchar"`       // Foreign key for Product
	Amount      float64    `gorm:"not null"`                    // Monetary amount
	Qty         int32      `gorm:"not null"`                    // Quantity of products
	CreatedAt   time.Time  `gorm:"not null;autoCreateTime"`     // Automatically created timestamp
	UpdatedAt   time.Time  `gorm:"not null;autoUpdateTime"`     // Automatically updated timestamp
	DiscardedAt *time.Time `gorm:"index"`                       // Soft delete timestamp
}

func (PaymentItem) TableName() string {
	return "paymentitem"
}
