package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getPathParameterFromRequest(req *http.Request, variableName string) (string, bool) {
	vars := mux.Vars(req)
	value, ok := vars[variableName]
	return value, ok
}
