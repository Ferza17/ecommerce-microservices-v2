package response

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

// Response helpers
type response struct {
	Error   string          `json:"error"`
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response{
		Error:   err.Error(),
		Message: message,
		Code:    statusCode,
	})
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(response{
		Data: func() json.RawMessage {
			r := []byte("null")
			if data == nil {
				return json.RawMessage(r)
			}

			if msg, ok := data.(proto.Message); ok {
				marshaler := &protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				}
				jsonBytes, err := marshaler.Marshal(msg)
				if err != nil {
					return json.RawMessage(r)
				}
				return json.RawMessage(jsonBytes)
			}

			jsonBytes, err := json.Marshal(data)
			if err != nil {
				return json.RawMessage(r)
			}
			return json.RawMessage(jsonBytes)
		}(),
		Message: "Success",
		Code:    statusCode,
	})
}
