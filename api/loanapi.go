package api

import (
	"net/http"
)

// InitAPIRoutes : Initialize api
func InitAPIRoutes(rt *http.ServeMux) {
	rt.HandleFunc("/generate-plan", generateLoanPlan)
}

// generateLoanPlan : Http post get
func generateLoanPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		return
	}
	jsonErrorResponse("Only HTTP.POST is supported", http.StatusMethodNotAllowed, w)
}
