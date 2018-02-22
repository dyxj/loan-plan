package money

import (
	"math/big"
	"strconv"
)

// Cent2DollarStr : Returns dollars given cents
func Cent2DollarStr(cents int64) string {
	strC := strconv.FormatInt(cents, 10)
	lstrC := len(strC)
	if lstrC > 2 {
		strC = strC[:lstrC-2] + "." + strC[lstrC-2:]
	} else if lstrC > 1 {
		strC = "0." + strC
	} else {
		strC = "0.0" + strC
	}
	return strC
}

// RoundCent : Rounds to closest cent
func RoundCent(cents *big.Float) int64 {
	centF := new(big.Float).Add(cents, big.NewFloat(0.5))
	centI, _ := centF.Int64()
	return centI
}
