package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *userService) FindUserByEmailAndPassword(ctx context.Context, requestId string, req *pb.FindUserByEmailAndPasswordRequest) (*pb.User, error) {
	result, err := s.cb.Execute(func() (interface{}, error) {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			enum.XRequestIDHeader.String(): requestId,
		}))
		resp, err := s.svc.FindUserByEmailAndPassword(ctx, req)
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				if st.Code() == codes.NotFound {
					return nil, gobreaker.ErrOpenState
				}
			}
			return nil, err
		}
		return resp, nil
	})

	if err != nil {
		if err == gobreaker.ErrOpenState {
			s.logger.Error(fmt.Sprintf("Circuit Breaker for User Service is open. Request Failed: %v\n", err))
			return nil, status.Errorf(codes.Unavailable, "User Service is currently unavailable")
		}
		if err == gobreaker.ErrTooManyRequests {
			s.logger.Error(fmt.Sprintf("Circuit Breaker for User Service in half-open mode and too many request: %v\n", err))
			return nil, status.Errorf(codes.Unavailable, "User Service is busy, please try again later")
		}
		return nil, fmt.Errorf("failed to call User Service: %w", err)
	}

	return result.(*pb.User), nil
}
