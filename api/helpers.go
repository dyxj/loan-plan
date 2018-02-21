package api

import (
	"fmt"
	"net/http"
)

func jsonErrorResponse(errMsg string, statusCode int, w http.ResponseWriter) {
	jsStr := fmt.Sprintf("{\"error\":\"%s\"}", errMsg)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(jsStr))
}
