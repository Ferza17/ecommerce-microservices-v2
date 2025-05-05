package rabbitmq

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/postgresql"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/product-service/pkg"
)

func NewLogger(logger pkg.IZapLogger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewPostgresConnector(pgconn postgresql.IPostgreSQLInfrastructure) Option {
	return func(s *Server) {
		s.postgresqlConnector = pgconn
	}
}

func NewRabbitMQInfrastructure(mq rabbitmq.IRabbitMQInfrastructure) Option {
	return func(s *Server) {
		s.rabbitmqInfrastructure = mq
	}
}
