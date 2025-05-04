package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/enum"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"time"
)

type PostgresqlConnector struct {
	GormDB *gorm.DB
	SqlDB  *sql.DB
}

func NewPostgresqlConnector() *PostgresqlConnector {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get().PostgresHost,
		config.Get().PostgresPort,
		config.Get().PostgresUsername,
		config.Get().PostgresPassword,
		config.Get().PostgresDatabaseName,
		config.Get().PostgresSSLMode)

	sqldb, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("could not open sql %s:", err))
	}

	log.Println("SQLDB registered")

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
		log.Fatal(fmt.Sprintf("could not open dao %s:", err))
	}

	gormSqlDB, err := gormdb.DB()
	if err != nil {
		log.Fatal(fmt.Sprintf("could not get dao DB %s:", err))
	}

	if err = gormSqlDB.Ping(); err != nil {
		log.Fatal(fmt.Sprintf("could not dao db ping %s:", err))
	}
	gormSqlDB.SetMaxOpenConns(10)
	gormSqlDB.SetMaxIdleConns(5)
	gormSqlDB.SetConnMaxIdleTime(300 * time.Second)
	gormSqlDB.SetConnMaxLifetime(time.Duration(300 * time.Second))

	log.Println("GORM registered")

	return &PostgresqlConnector{
		GormDB: gormdb,
		SqlDB:  sqldb,
	}
}

func (p *PostgresqlConnector) Close() error {
	return p.SqlDB.Close()
}
