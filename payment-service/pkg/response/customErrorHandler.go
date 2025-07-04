package response

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	// Extract gRPC status from error
	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Unknown, err.Error())
	}

	httpStatusCode := runtime.HTTPStatusFromCode(st.Code())

	errorResp := response{
		Error:   st.Code().String(),
		Message: st.Message(),
		Code:    httpStatusCode,
		Data:    json.RawMessage(`null`),
	}

	switch st.Code() {
	case codes.InvalidArgument:
		errorResp = response{
			Code:    400,
			Message: fmt.Sprintf("Invalid request parameters : %s", st.Message()),
		}
	case codes.NotFound:
		errorResp = response{
			Code:    404,
			Message: fmt.Sprintf("Resource not found : %s", st.Message()),
		}
	case codes.PermissionDenied:
		errorResp = response{
			Code:    403,
			Message: fmt.Sprintf("Permission denied : %s", st.Message()),
		}
	case codes.Unauthenticated:
		errorResp = response{
			Code:    401,
			Message: fmt.Sprintf("Authentication failed : %s", st.Message()),
		}
	case codes.Internal:
		errorResp = response{
			Code:    500,
			Message: "Internal server error",
		}
	default:
		errorResp = response{
			Code:    500,
			Message: "Unknown error",
		}
	}

	// Marshal response to JSON
	buf, marshalErr := marshaler.Marshal(errorResp)
	if marshalErr != nil {
		// Fallback to basic error response if marshaling fails
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"INTERNAL","message":"Internal server error","code":500,"data":null}`))
		return
	}

	// Set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write(buf)
}
