package dto

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *Response) WriteResponse(w http.ResponseWriter, statusCode int, message string) error {
	r.Message = message
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
