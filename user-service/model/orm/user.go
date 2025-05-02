package orm

import (
	"time"
)

type User struct {
	ID          string     `gorm:"column:id;primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Email       string     `gorm:"column:email" json:"email"`
	Password    string     `gorm:"column:password" json:"password"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DiscardedAt *time.Time `gorm:"column:discarded_at" json:"discarded_at"`
}
