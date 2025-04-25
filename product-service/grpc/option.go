package grpc

import "go.uber.org/zap"

func NewLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}
