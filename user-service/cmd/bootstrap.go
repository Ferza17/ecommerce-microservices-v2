package cmd

import (
	"context"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	amqp "github.com/ferza17/ecommerce-microservices-v2/user-service/server/amqp"
	grpc "github.com/ferza17/ecommerce-microservices-v2/user-service/server/grpc"
	"log"
)

type Bootstrap struct {
}

var (
	logger      pkg.IZapLogger
	pgsqlConn   *connector.PostgresqlConnector
	grpcServer  *grpc.Server
	mongoDBConn *connector.MongodbConnector
	amqpServer  *amqp.Server
)

func init() {
	config.SetConfig(".")
	logger = pkg.NewZapLogger()
	pgsqlConn = connector.NewPostgresqlConnector()
	mongoDBConn = connector.NewMongodbConnector()
	grpcServer = NewGrpcServer()
	amqpServer = NewAmqpServer()
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

func NewGrpcServer() (srv *grpc.Server) {
	return grpc.NewServer(
		config.Get().RpcHost,
		config.Get().RpcPort,
		grpc.NewLogger(logger),
		grpc.NewPostgresConnector(pgsqlConn),
		grpc.NewMongoDBConnector(mongoDBConn),
	)
}

func NewAmqpServer() (srv *amqp.Server) {
	return amqp.NewServer(
		amqp.NewLogger(logger),
		amqp.NewPostgresConnector(pgsqlConn),
		amqp.NewMongoDBConnector(mongoDBConn),
	)
}
