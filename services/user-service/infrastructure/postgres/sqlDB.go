package postgres

import "database/sql"

func (p *postgresSQL) SqlDB() *sql.DB {
	return p.sqlDB
}
