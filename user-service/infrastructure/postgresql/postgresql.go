package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type (
	IPostgreSQLInfrastructure interface {
		Close() error
		GormDB() *gorm.DB
		SqlDB() *sql.DB
	}

	PostgreSQLInfrastructure struct {
		gormDB *gorm.DB
		sqlDB  *sql.DB
		logger pkg.IZapLogger
	}
)

func NewPostgresqlInfrastructure(logger pkg.IZapLogger) IPostgreSQLInfrastructure {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get().PostgresHost,
		config.Get().PostgresPort,
		config.Get().PostgresUsername,
		config.Get().PostgresPassword,
		config.Get().PostgresDatabaseName,
		config.Get().PostgresSSLMode)

	sqldb, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	gormConfig := gorm.Config{}
	if config.Get().Env == enum.CONFIG_ENV_LOCAL {
		gormConfig = gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
		}
	}

	pgDialect := postgres.New(postgres.Config{
		Conn: sqldb,
	})

	gormdb, err := gorm.Open(pgDialect, &gormConfig)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	gormSqlDB, err := gormdb.DB()
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	if err = gormSqlDB.Ping(); err != nil {
		logger.Error(fmt.Sprintf("failed to connect to postgres: %v", err))
	}
	gormSqlDB.SetMaxOpenConns(10)
	gormSqlDB.SetMaxIdleConns(5)
	gormSqlDB.SetConnMaxIdleTime(300 * time.Second)
	gormSqlDB.SetConnMaxLifetime(time.Duration(300 * time.Second))

	return &PostgreSQLInfrastructure{
		gormDB: gormdb,
		sqlDB:  sqldb,
	}
}

func (p *PostgreSQLInfrastructure) Close() error {
	return p.sqlDB.Close()
}

func (p *PostgreSQLInfrastructure) GormDB() *gorm.DB {
	return p.gormDB
}

func (p *PostgreSQLInfrastructure) SqlDB() *sql.DB {
	return p.sqlDB
}
