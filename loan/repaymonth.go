package loan

import (
	"time"
)

// RepayMonth : all values are in cents
type RepayMonth struct {
	BPAmount      int64     `json:"borrowerPaymentAmount"`
	Date          time.Time `json:"date"`
	IOutPrincipal int64     `json:"initialOutstandingPrincipal"`
	Interest      int64     `json:"interest"`
	Principal     int64     `json:"principal"`
	ROutPrincipal int64     `json:"remainingOutstandingPrincipal"`
}

// RepayMonthDollars : all values are in dollars
type RepayMonthDollars struct {
	BPAmount      float64   `json:"borrowerPaymentAmount"`
	Date          time.Time `json:"date"`
	IOutPrincipal float64   `json:"initialOutstandingPrincipal"`
	Interest      float64   `json:"interest"`
	Principal     float64   `json:"principal"`
	ROutPrincipal float64   `json:"remainingOutstandingPrincipal"`
}

// Convert2Dollars : Convert for cents type to dollar type
func (rm *RepayMonth) Convert2Dollars() *RepayMonthDollars {
	rmd := &RepayMonthDollars{
		BPAmount:      float64(rm.BPAmount) / 100.00,
		Date:          rm.Date,
		IOutPrincipal: float64(rm.IOutPrincipal) / 100.00,
		Interest:      float64(rm.Interest) / 100.00,
		Principal:     float64(rm.Principal) / 100.00,
		ROutPrincipal: float64(rm.ROutPrincipal) / 100.00,
	}
	return rmd
}
