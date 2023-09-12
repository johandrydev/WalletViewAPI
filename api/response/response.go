package response

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Create(w http.ResponseWriter, message string, status int, data any) {
	response := BaseResponse{
		Message: message,
	}
	if data != nil {
		response.Data = data
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonResponse)
}
