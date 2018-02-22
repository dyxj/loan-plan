package money_test

import (
	"github.com/dyxj/loan-plan/money"
	"math/big"
	"testing"
)

func TestCent2DollarStr(t *testing.T) {
	// test cases
	ttc := []struct {
		name      string
		cents     int64
		dollarstr string
	}{
		{"1 cent", 1, "0.01"},
		{"10 cent", 10, "0.10"},
		{"100 cent", 100, "1.00"},
		{"1000 cent", 1000, "10.00"},
		{"1234 cent", 1234, "12.34"},
		{"12345 cent", 12345, "123.45"},
	}

	// Sub test
	for _, tc := range ttc {
		t.Run(tc.name, func(t *testing.T) {
			ds := money.Cent2DollarStr(tc.cents)
			if ds != tc.dollarstr {
				t.Fatalf("expected %v for %v got: %v", tc.dollarstr, tc.cents, ds)
			}
		})
	}
}

func TestRoundCent(t *testing.T) {
	// test cases
	ttc := []struct {
		name   string
		cents  *big.Float
		rcents int64
	}{
		{"1 cent", big.NewFloat(1), 1},
		{"1.5 cent", big.NewFloat(1.5), 2},
		{"1.6 cent", big.NewFloat(1.6), 2},
		{"1.0 cent", big.NewFloat(1.0), 1},
		{"1.0 cent", big.NewFloat(1.167), 1},
		{"1.0 cent", big.NewFloat(1.678), 2},
	}

	// Sub test
	for _, tc := range ttc {
		t.Run(tc.name, func(t *testing.T) {
			rc := money.RoundCent(tc.cents)
			if rc != tc.rcents {
				t.Fatalf("expected %v got %v", tc.rcents, rc)
			}
		})
	}
}
