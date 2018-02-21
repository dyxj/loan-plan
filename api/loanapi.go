package api

import (
	"encoding/json"
	"github.com/dyxj/loan-plan/loan"
	"log"
	"net/http"
	"strconv"
)

// InitAPIRoutes : Initialize api
func InitAPIRoutes(rt *http.ServeMux) {
	rt.HandleFunc("/generate-plan", generateLoanPlan)
}

// generateLoanPlan : Http post get
func generateLoanPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonErrorResponse("Only HTTP.POST is supported", http.StatusMethodNotAllowed, w)
		return
	}

	d := json.NewDecoder(r.Body)
	var preq loan.PlanReq
	err := d.Decode(&preq)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Unexpected error occured when decoding request: "+err.Error(), http.StatusInternalServerError, w)
		return
	}

	// Get total loan amount as float
	tlaF, err := strconv.ParseFloat(preq.LoanAmount, 64)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid loan amount.", http.StatusInternalServerError, w)
		return
	}
	// Convert to int64, cents
	tlaICents := int64(tlaF * 100.00)

	// Get nominal interest rate
	nir, err := strconv.ParseFloat(preq.NominalRate, 64)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid nominal rate.", http.StatusInternalServerError, w)
		return
	}

	// Get payment plan
	slRM, err := loan.GenPlan(tlaICents, nir, preq.Duration, preq.StartDate)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Error occured generating payment plan", http.StatusInternalServerError, w)
		return
	}

	// Convert to dollars
	slRMD := make([]*loan.RepayMonthDollars, len(slRM))
	for i, v := range slRM {
		slRMD[i] = v.Convert2Dollars()
	}

	// Send array as json
	jsonResponse(slRMD, http.StatusOK, w)
}
