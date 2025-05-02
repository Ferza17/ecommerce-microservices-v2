package orm

import "time"

type Product struct {
	ID          string     `gorm:"column:id;primaryKey"`
	Name        string     `gorm:"column:name"`
	Price       float64    `gorm:"column:price"`
	Stock       int64      `gorm:"column:stock"`
	Description string     `gorm:"column:description"`
	Image       string     `gorm:"column:image"`
	Uom         string     `gorm:"column:uom"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DiscardedAt *time.Time `gorm:"column:discarded_at"`
}
