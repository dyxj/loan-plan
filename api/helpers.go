package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ----------------- JSON Response -----------------
func jsonResponse(resp interface{}, statusCode int, w http.ResponseWriter) {
	js, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error occured creating JSON response: %v\n", err)
		jsonErrorResponse("Error occured creating JSON response",
			http.StatusInternalServerError, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

func jsonErrorResponse(errMsg string, statusCode int, w http.ResponseWriter) {
	jsStr := fmt.Sprintf("{\"error\":\"%s\"}", errMsg)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(jsStr))
}

// ----------------- End JSON Response -----------------
