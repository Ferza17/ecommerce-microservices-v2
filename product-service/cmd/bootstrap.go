// Bootstaping The Service

package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	pgsql "github.com/ferza17/ecommerce-microservices-v2/product-service/connector/postgresql"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	logger    *zap.Logger
	pgsqlConn *pgsql.PostgresqlConnector
)

func init() {
	config.SetConfig(".")
	logger = NewLogger()
	pgsqlConn = pgsql.NewPostgresqlConnector()
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
