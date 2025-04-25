package helper

import (
	"github.com/ferza17/ecommerce-microservices-v2/product-service/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RpcErrorHandler(errArgs error) error {
	var (
		mapError = map[string]error{
			// Internal Error
			utils.ErrInternalServerError.Error(): status.Error(codes.Internal, utils.ErrInternalServerError.Error()),
			utils.ErrQueryTxInsert.Error():       status.Error(codes.Internal, utils.ErrQueryTxInsert.Error()),
			utils.ErrQueryTxRollback.Error():     status.Error(codes.Internal, utils.ErrQueryTxRollback.Error()),
			utils.ErrQueryTxCommit.Error():       status.Error(codes.Internal, utils.ErrQueryTxRollback.Error()),
			utils.ErrQueryTxBegin.Error():        status.Error(codes.Internal, utils.ErrQueryTxBegin.Error()),
			utils.ErrQueryTxUpdate.Error():       status.Error(codes.Internal, utils.ErrQueryTxUpdate.Error()),
			utils.ErrQueryRead.Error():           status.Error(codes.Internal, utils.ErrQueryRead.Error()),

			// Bad Request
			utils.ErrBadRequest.Error(): status.Error(codes.InvalidArgument, utils.ErrBadRequest.Error()),

			// Not Found
			utils.ErrNotFound.Error(): status.Error(codes.NotFound, utils.ErrNotFound.Error()),
		}
	)
	if _, ok := mapError[errArgs.Error()]; !ok {
		return status.Error(codes.Unknown, "Unknown")
	}
	return mapError[errArgs.Error()]
}
