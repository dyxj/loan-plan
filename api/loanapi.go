package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dyxj/loan-plan/money"

	"github.com/dyxj/loan-plan/loan"
)

// InitAPIRoutes : Initialize api
func InitAPIRoutes(rt *http.ServeMux) {
	rt.HandleFunc("/generate-plan", generateLoanPlan)
	rt.HandleFunc("/calc-annuity", calcAnnuity)
}

// calcAnnuity: Http post to generate annuity
func calcAnnuity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonErrorResponse("Only HTTP POST is supported", http.StatusMethodNotAllowed, w)
		return
	}

	d := json.NewDecoder(r.Body)
	var preq loan.PlanReq
	err := d.Decode(&preq)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Unexpected error occured when decoding request: "+err.Error(), http.StatusBadRequest, w)
		return
	}

	// Get total loan amount as float
	tlaF, err := strconv.ParseFloat(preq.LoanAmount, 64)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid loan amount.", http.StatusBadRequest, w)
		return
	}
	// Convert to int64, cents
	tlaCents := int64(tlaF * 100.00)

	// Get nominal interest rate
	nir, err := strconv.ParseFloat(preq.NominalRate, 64)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid nominal rate.", http.StatusBadRequest, w)
		return
	}

	// Rate per period in %
	rpp := nir * 30 / 360

	// Calculate Annuity
	a := loan.CalcAnnuity(tlaCents, rpp, preq.Duration)
	aDollars := money.Cent2DollarStr(a)

	apiResp := struct {
		Annuity string `json:"annuity"`
	}{
		Annuity: aDollars,
	}

	// Send annuity struct as json
	jsonResponse(apiResp, http.StatusOK, w)
}

// generateLoanPlan : Http post to generate loan plan
func generateLoanPlan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		jsonErrorResponse("Only HTTP POST is supported", http.StatusMethodNotAllowed, w)
		return
	}

	d := json.NewDecoder(r.Body)
	var preq loan.PlanReq
	err := d.Decode(&preq)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Unexpected error occured when decoding request: "+err.Error(), http.StatusBadRequest, w)
		return
	}

	// Get total loan amount as float
	tlaF, err := strconv.ParseFloat(preq.LoanAmount, 64)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid loan amount.", http.StatusBadRequest, w)
		return
	}
	// Convert to int64, cents
	tlaICents := int64(tlaF * 100.00)

	// Get nominal interest rate
	nir, err := strconv.ParseFloat(preq.NominalRate, 64)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid nominal rate.", http.StatusBadRequest, w)
		return
	}

	// Get start date
	sd, err := time.Parse("2006-01-02", preq.StartDate)
	if err != nil {
		log.Println(err)
		jsonErrorResponse("Invalid date format.", http.StatusBadRequest, w)
		return
	}

	// Generate payment plan
	slRM, err := loan.GenPlan(tlaICents, nir, preq.Duration, sd)
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
