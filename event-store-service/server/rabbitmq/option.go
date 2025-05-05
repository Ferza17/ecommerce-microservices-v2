package rabbitmq

import (
	mongoDBInfrastructure "github.com/ferza17/ecommerce-microservices-v2/event-store-service/infrastructure/mongodb"
	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/pkg"
)

func NewLogger(logger pkg.IZapLogger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewMongoDBInfrastructure(mongoDBInfrastructure mongoDBInfrastructure.IMongoDBInfrastructure) Option {
	return func(s *Server) {
		s.mongoDBInfrastructure = mongoDBInfrastructure
	}
}
