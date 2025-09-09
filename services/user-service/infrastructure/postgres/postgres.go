package postgres

import (
	"database/sql"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type (
	IPostgresSQL interface {
		GormDB() *gorm.DB
		SqlDB() *sql.DB
		Close() error
	}

	postgresSQL struct {
		gormDB *gorm.DB
		sqlDB  *sql.DB
		logger logger.IZapLogger
	}
)

var Set = wire.NewSet(NewPostgresqlInfrastructure)

func NewPostgresqlInfrastructure(logger logger.IZapLogger) IPostgresSQL {
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
		panic(err)
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

	return &postgresSQL{
		gormDB: gormdb,
		sqlDB:  sqldb,
	}
}

func (p *postgresSQL) Close() error {
	return p.sqlDB.Close()
}
