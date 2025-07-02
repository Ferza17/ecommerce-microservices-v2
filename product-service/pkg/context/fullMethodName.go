package context

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetFullMethodNameFromContext(ctx context.Context) (string, error) {
	method := grpc.ServerTransportStreamFromContext(ctx).Method()
	if method != "" {
		return method, nil
	}
	if method, _ = runtime.RPCMethod(ctx); method != "" {
		return method, status.Error(codes.Unimplemented, "unimplemented method")
	}
	return method, nil
}
