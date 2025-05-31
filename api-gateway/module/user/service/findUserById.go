package service

import (
	"context"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *userService) FindUserById(ctx context.Context, requestId string, req *userRpc.FindUserByIdRequest) (*userRpc.User, error) {
	result, err := s.cb.Execute(func() (interface{}, error) {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			enum.XRequestIDHeader.String(): requestId,
		}))
		resp, err := s.svc.FindUserById(ctx, req)
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
		s.logger.Error(fmt.Sprintf("Error Breaker %v", err))
		return nil, err
	}
	return result.(*userRpc.User), err
}
