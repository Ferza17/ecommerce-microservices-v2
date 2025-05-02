package orm

import (
	"time"
)

type User struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name" validate:"required"`
	Email       string     `gorm:"column:email" json:"email" validate:"required,email"`
	Password    string     `gorm:"column:password" json:"password" validate:"required"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DiscardedAt *time.Time `gorm:"column:discarded_at"`
}
