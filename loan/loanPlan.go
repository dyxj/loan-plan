package loan

import (
	"github.com/dyxj/loan-plan/money"
	"math"
	"math/big"
	"time"
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
	// Initial Outstanding
	iop := tlaCents

	// Calculate Annuity
	a := calcAnnuity(tlaCents, rpp, dur)

	slRM := make([]*RepayMonth, dur)

	for i := 0; i < dur; i++ {
		// Calculate Interest
		it := calcInterest(iop, nir)

		//  Calculate Principal
		p := calcPrincipal(iop, a, it)

		// Calculate BPA
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

func calcAnnuity(pv int64, rpp float64, periods int) int64 {
	bpv := big.NewFloat(float64(pv))
	brpp := big.NewFloat(rpp / 100)
	annuity := new(big.Float).Mul(bpv, brpp)
	adenom := math.Pow(1+rpp/100, float64(-1*periods))
	bAdenom := new(big.Float).Sub(big.NewFloat(1), big.NewFloat(adenom))
	annuity = annuity.Quo(annuity, bAdenom)
	annuity = money.RoundCent(annuity)
	intC, _ := annuity.Int64()
	return intC
}

// outstandingAmt, Outstanding Amount in cents.
// nir, Nominal Interest Rate in percentage. ie: 5.00%.
func calcInterest(oaCents int64, nir float64) int64 {
	btla := big.NewFloat(float64(oaCents))
	bdayRatio := big.NewFloat(dayRatio)
	bnir := big.NewFloat(nir / 100)
	interest := new(big.Float).Mul(btla, bdayRatio)
	interest.Mul(interest, bnir)
	interest = money.RoundCent(interest)
	intC, _ := interest.Int64()
	return intC
}

func calcPrincipal(oaCents, a, i int64) int64 {
	principal := a - i
	if principal > oaCents {
		return oaCents
	}
	return principal
}
