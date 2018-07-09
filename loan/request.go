package loan

// PlanReq : Loan plan request
type PlanReq struct {
	LoanAmount  string `json:"loanAmount"`
	NominalRate string `json:"nominalRate"`
	Duration    int    `json:"duration"`
	StartDate   string `json:"startDate,omitempty"`
}
