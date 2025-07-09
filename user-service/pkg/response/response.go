package response

import (
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/common/response"
	"go.opentelemetry.io/otel/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
	"net/http"
)

var marshalOptions = protojson.MarshalOptions{
	Multiline:       true,
	Indent:          "  ",
	UseProtoNames:   false,
	EmitUnpopulated: false,
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	marshal, err := marshalOptions.Marshal(&pb.Response{
		Error:   err.Error(),
		Message: message,
		Code:    int32(statusCode),
	})
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(marshal)
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data *structpb.Struct) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	marshal, err := marshalOptions.Marshal(&pb.Response{
		Error:   "",
		Message: codes.Ok.String(),
		Code:    int32(statusCode),
		Data:    data,
	})
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(marshal)
}
