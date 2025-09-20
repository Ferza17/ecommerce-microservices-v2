package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/enum"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/pkg/logger"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type (
	PostgresSQL struct {
		GormDB *gorm.DB
		SqlDB  *sql.DB
		logger logger.IZapLogger
	}
)

// Set is a Wire provider set for PostgresSQL dependencies
var Set = wire.NewSet(
	NewPostgresqlInfrastructure,
)

func NewPostgresqlInfrastructure(logger logger.IZapLogger) *PostgresSQL {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get().DatabasePostgres.Host,
		config.Get().DatabasePostgres.Port,
		config.Get().DatabasePostgres.Username,
		config.Get().DatabasePostgres.Password,
		config.Get().DatabasePostgres.Database,
		config.Get().DatabasePostgres.SSLMode)

	sqldb, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to connect to postgres: %v", err))
	}

	gormConfig := gorm.Config{}
	if config.Get().Env == enum.CONFIG_ENV_LOCAL {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Info)
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

	return &PostgresSQL{
		GormDB: gormdb,
		SqlDB:  sqldb,
	}
}
