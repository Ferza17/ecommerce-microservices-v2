package orm

import "time"

type Product struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Price       float64    `gorm:"column:price" json:"price"`
	Stock       int64      `gorm:"column:stock" json:"stock"`
	Description string     `gorm:"column:description" json:"description"`
	Image       string     `gorm:"column:image" json:"image"`
	Uom         string     `gorm:"column:uom" json:"uom"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DiscardedAt *time.Time `gorm:"column:discarded_at" json:"discarded_at,omitempty"`
}
