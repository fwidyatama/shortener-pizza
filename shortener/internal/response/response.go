package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data"`
	Error  *Error      `json:"error,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	resp := Response{}
	resp.Status = 200
	resp.Error = nil
	resp.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}

func ErrorResponse(w http.ResponseWriter, message string, statusCode int) {

	resp := Response{}
	resp.Status = statusCode
	resp.Error = &Error{}
	resp.Error.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)

}
