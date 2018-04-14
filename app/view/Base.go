package view

import "net/http"

type ApiResponse struct {
	Headers map[string]string `json:"-"`
	Status  int
	Msg     interface{}       `json:"Msg,omitempty"`
}

func Ok(data interface{}) ApiResponse {
	return ApiResponse{Status: http.StatusOK, Msg: data}
}