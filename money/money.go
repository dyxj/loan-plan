package money

import (
	"math/big"
	"strconv"
)

// Cent2Dollar : Returns dollars given cents
func Cent2Dollar(cents int64) float64 {
	return float64(cents) / 100.00
}

// Cent2DollarStr : Returns dollars given cents
func Cent2DollarStr(cents int64) string {
	return strconv.FormatFloat(Cent2Dollar(cents), 'f', 2, 64)
}

// RoundCent :
func RoundCent(cents *big.Float) *big.Float {
	return new(big.Float).Add(cents, big.NewFloat(0.5))
}
