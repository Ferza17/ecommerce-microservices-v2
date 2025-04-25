// Bootstaping The Service

package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/enum"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"log"
	"time"
)

var (
	logger *zap.Logger
	gormDB *gorm.DB
	sqlDB  *sql.DB
)

func init() {
	config.SetConfig(".")
	logger = NewLogger()
	sqlDB = NewSQLDB()
	gormDB = NewPostgresSQlClient()
}

func Shutdown(ctx context.Context) (err error) {
	//cassandraSession.Close()
	//if err = postgresSQlClient.Close(); err != nil {
	//	return
	//}
	//if err = redisClient.Close(); err != nil {
	//	return
	//}
	//if err = rabbitMQConnection.Close(); err != nil {
	//	return
	//}
	return
}

func NewSQLDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get().PostgresHost,
		config.Get().PostgresPort,
		config.Get().PostgresUsername,
		config.Get().PostgresPassword,
		config.Get().PostgresDatabaseName,
		config.Get().PostgresSSLMode)

	sqldb, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Fatal(fmt.Sprintf("could not open sql %s:", err))
	}
	log.Println("SQLDB registered")

	return sqldb
}

func NewPostgresSQlClient() *gorm.DB {

	gormConfig := gorm.Config{}
	if config.Get().Env == enum.CONFIG_ENV_LOCAL {
		gormConfig = gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
		}
	}

	pgDialect := postgres.New(postgres.Config{
		Conn: sqlDB,
	})

	gormdb, err := gorm.Open(pgDialect, &gormConfig)
	if err != nil {
		logger.Fatal(fmt.Sprintf("could not open gorm %s:", err))
	}

	gormSqlDB, err := gormdb.DB()
	if err != nil {
		logger.Fatal(fmt.Sprintf("could not get gorm DB %s:", err))
	}

	if err = gormSqlDB.Ping(); err != nil {
		logger.Fatal(fmt.Sprintf("could not gorm db ping %s:", err))
	}
	gormSqlDB.SetMaxOpenConns(10)
	gormSqlDB.SetMaxIdleConns(5)
	gormSqlDB.SetConnMaxIdleTime(300 * time.Second)
	gormSqlDB.SetConnMaxLifetime(time.Duration(300 * time.Second))

	log.Println("GORM registered")
	return gormdb
}

func NewLogger() (logger *zap.Logger) {
	var err error
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	if logger, err = logConfig.Build(); err != nil {
		log.Fatalf("error when register logger: %v\n", err)
	}
	log.Println("LOGGER registered")
	return
}
