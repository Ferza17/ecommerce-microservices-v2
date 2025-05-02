// Bootstaping The Service

package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	logger      *zap.Logger
	pgsqlConn   *connector.PostgresqlConnector
	grpcServer  *grpc.Server
	mongoDBConn *connector.MongodbConnector
)

func init() {
	config.SetConfig(".")
	logger = NewLogger()
	pgsqlConn = connector.NewPostgresqlConnector()
	mongoDBConn = connector.NewMongodbConnector()
	grpcServer = NewGrpcServer()
}

func Shutdown(ctx context.Context) (err error) {
	grpcServer.GracefulStop()

	if err = pgsqlConn.Close(); err != nil {
		return err
	}
	if err = mongoDBConn.Close(ctx); err != nil {
		return err
	}

	log.Println("Shutdown...")
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

func NewGrpcServer() (srv *grpc.Server) {
	return grpc.NewServer(
		config.Get().RpcHost,
		config.Get().RpcPort,
		grpc.NewLogger(logger),
		grpc.NewPostgresConnector(pgsqlConn),
	)
}
