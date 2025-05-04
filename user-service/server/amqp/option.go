package amqp

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
)

func NewLogger(logger pkg.IZapLogger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewPostgresConnector(pgconn *infrastructure.PostgresqlConnector) Option {
	return func(s *Server) {
		s.postgresqlConnector = pgconn
	}
}

func NewMongoDBConnector(mongodbConnector *infrastructure.MongodbConnector) Option {
	return func(s *Server) {
		s.mongoDBConnector = mongodbConnector
	}
}
