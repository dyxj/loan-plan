package loan

import (
	"time"
)

// PlanReq : Loan plan request
type PlanReq struct {
	LoanAmount  string    `json:"loanAmount"`
	NominalRate string    `json:"nominalRate"`
	Duration    int       `json:"duration"`
	StartDate   time.Time `json:"startDate"`
}
