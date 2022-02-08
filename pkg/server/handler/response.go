package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func makeResponse(res http.ResponseWriter, statusCode int, body interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(body)
}
