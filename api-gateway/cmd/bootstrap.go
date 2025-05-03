package cmd

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
	"log"
)

var (
	logger      pkg.IZapLogger
	mongoDBConn connector.IMongodbConnector
	amqpConn    connector.IAmqpConnector
)

func init() {
	logger = pkg.NewZapLogger()
	mongoDBConn = connector.NewMongodbConnector()
	amqpConn = connector.NewAmqpConnector()
}

func Shutdown(ctx context.Context) (err error) {
	if err = amqpConn.Close(); err != nil {
		logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	if err = mongoDBConn.Close(ctx); err != nil {
		logger.Error(fmt.Sprintf("Failed to close a connection: %v", err))
		return err
	}

	log.Println("Shutdown...")
	return
}

//func NewGraphQLServer()  (srv *graphql.Server){}
