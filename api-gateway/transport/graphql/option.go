package graphql

import bootstrap2 "github.com/ferza17/ecommerce-microservices-v2/api-gateway/bootstrap"

func NewBootstrap(c *bootstrap2.Bootstrap) Option {
	return func(s *GraphQLTransport) {
		s.bootstrap = c
	}
}
