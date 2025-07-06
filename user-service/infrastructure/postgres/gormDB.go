package postgres

import "gorm.io/gorm"

func (p *postgresSQL) GormDB() *gorm.DB {
	return p.gormDB
}
