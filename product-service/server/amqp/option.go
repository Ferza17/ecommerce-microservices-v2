package amqp

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

func NewLogger(logger pkg.IZapLogger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewPostgresConnector(pgconn *connector.PostgresqlConnector) Option {
	return func(s *Server) {
		s.postgresqlConnector = pgconn
	}
}

func NewMongoDBConnector(mongodbConnector *connector.MongodbConnector) Option {
	return func(s *Server) {
		s.mongoDBConnector = mongodbConnector
	}
}
