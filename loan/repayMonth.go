package loan

import (
	"github.com/dyxj/loan-plan/money"
	// "strconv"
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
	BPAmount      string    `json:"borrowerPaymentAmount"`
	Date          time.Time `json:"date"`
	IOutPrincipal string    `json:"initialOutstandingPrincipal"`
	Interest      string    `json:"interest"`
	Principal     string    `json:"principal"`
	ROutPrincipal string    `json:"remainingOutstandingPrincipal"`
}

// Convert2Dollars : Convert for cents type to dollar type
func (rm *RepayMonth) Convert2Dollars() *RepayMonthDollars {
	rmd := &RepayMonthDollars{
		BPAmount:      money.Cent2DollarStr(rm.BPAmount),
		Date:          rm.Date,
		IOutPrincipal: money.Cent2DollarStr(rm.IOutPrincipal),
		Interest:      money.Cent2DollarStr(rm.Interest),
		Principal:     money.Cent2DollarStr(rm.Principal),
		ROutPrincipal: money.Cent2DollarStr(rm.ROutPrincipal),
	}

	return rmd
}
