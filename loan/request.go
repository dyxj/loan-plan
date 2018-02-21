package loan

import (
	"time"
)

// PlanReq : Loan plan request
type PlanReq struct {
	LoanAmount  float64
	NominalRate float64
	Duration    int
	StartDate   time.Time
}
