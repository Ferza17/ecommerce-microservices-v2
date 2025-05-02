package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/connector"
	"go.uber.org/zap"
)

func NewLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewPostgresConnector(pgconn *connector.PostgresqlConnector) Option {
	return func(s *Server) {
		s.postgresqlConnector = pgconn
	}
}
