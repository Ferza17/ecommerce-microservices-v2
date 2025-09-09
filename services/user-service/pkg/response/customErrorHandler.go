package response

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"

	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/common/response"
)

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// Extract gRPC status from error
	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Unknown, err.Error())
	}

	httpStatusCode := runtime.HTTPStatusFromCode(st.Code())

	errorResp := pb.Response{
		Message: st.Message(),
		Status:  "error",
		Data:    nil,
	}

	switch st.Code() {
	case codes.InvalidArgument:
		errorResp = pb.Response{
			Message: fmt.Sprintf("Invalid request parameters : %s", st.Message()),
		}
	case codes.NotFound:
		errorResp = pb.Response{
			Message: fmt.Sprintf("Resource not found : %s", st.Message()),
		}
	case codes.PermissionDenied:
		errorResp = pb.Response{
			Message: fmt.Sprintf("Permission denied : %s", st.Message()),
		}
	case codes.Unauthenticated:
		errorResp = pb.Response{
			Message: fmt.Sprintf("Authentication failed : %s", st.Message()),
		}
	case codes.Internal:
		errorResp = pb.Response{
			Message: "Internal server error",
		}
	default:
		errorResp = pb.Response{
			Message: "Unknown error",
		}
	}

	WriteErrorResponse(w, httpStatusCode, errorResp.Message, err)
}
