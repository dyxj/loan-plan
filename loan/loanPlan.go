package loan

import (
	"math"
	"math/big"
	"time"

	"github.com/dyxj/loan-plan/money"
)

const (
	dayRatio float64 = 30.00 / 360.00
)

// GenPlan : return repayment plan given :-
// tla, Total Loan Amount in cents.
// nir, Nominal Interest Rate in percentage. ie: 5.00%.
// dur, Duration in months.
// sd, Start date.
func GenPlan(tlaCents int64, nir float64, dur int, sd time.Time) ([]*RepayMonth, error) {
	// Rate per period in %
	rpp := nir * 30 / 360
	// Initial Outstanding principal
	iop := tlaCents
	// Calculate Annuity
	a := CalcAnnuity(tlaCents, rpp, dur)

	slRM := make([]*RepayMonth, dur)
	for i := 0; i < dur; i++ {
		// Calculate Interest
		it := calcInterest(iop, nir)

		//  Calculate Principal
		p := calcPrincipal(iop, a, it)

		// Calculate Borrower Payment Amount(Annuity)
		bpa := p + it

		rm := &RepayMonth{
			BPAmount:      bpa,
			Date:          sd,
			IOutPrincipal: iop,
			Interest:      it,
			Principal:     p,
		}

		// Deduct Outstanding principal
		iop = iop - p
		rm.ROutPrincipal = iop

		// increment date by a month
		sd = sd.AddDate(0, 1, 0)
		slRM[i] = rm
	}

	return slRM, nil
}

// CalcAnnuity : calculates annuity for the duration of the perido
// pv, present value/total loan amount in cents
// rpp, rate per period in %
// periods in months
func CalcAnnuity(pv int64, rpp float64, periods int) int64 {
	bpv := big.NewFloat(float64(pv))
	brpp := big.NewFloat(rpp / 100)

	adenom := math.Pow(1+rpp/100, float64(-1*periods))
	bAdenom := new(big.Float).Sub(big.NewFloat(1), big.NewFloat(adenom))

	annuity := new(big.Float).Mul(bpv, brpp)
	annuity = annuity.Quo(annuity, bAdenom)
	annuityI := money.RoundCent(annuity)
	return annuityI
}

// calcInterest : calculates interest given
// oaCents, Outstanding Amount in cents.
// nir, Nominal Interest Rate in percentage. ie: 5.00%.
func calcInterest(oaCents int64, nir float64) int64 {
	btla := big.NewFloat(float64(oaCents))
	bdayRatio := big.NewFloat(dayRatio)
	bnir := big.NewFloat(nir / 100)

	interest := new(big.Float).Mul(btla, bdayRatio)
	interest.Mul(interest, bnir)
	interestI := money.RoundCent(interest)
	return interestI
}

// calcPrincipal : calculates principal amount
// oaCents, Outstanding Amount Principal in cents.
// a, Annuity in cents.
// i, Interest in cents.
// if calculated principal is more than outstanding principal amount
// it returns outstanding principal amount instead.
func calcPrincipal(oaCents, a, i int64) int64 {
	principal := a - i
	if principal > oaCents {
		return oaCents
	}
	return principal
}
