package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/connector"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg"
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
