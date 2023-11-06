package pdq

import (
	"testing"
)

func TestCalcLCOM(t *testing.T) {
	testCases := []struct {
		name                    string
		nonSharedAttributePairs int
		sharedAttributePairs    int
		expected                int
	}{
		{"More non-shared than shared attribute pairs", 10, 5, 5},
		{"More shared than non-shared attribute pairs", 5, 10, 0},
		{"Equal shared and non-shared attribute pairs", 5, 5, 0},
		{"Zero non-shared attribute pairs", 0, 5, 0},
		{"Zero shared attribute pairs", 5, 0, 5},
		{"Both pairs are zero", 0, 0, 0},
		{"Negative non-shared attribute pairs", -5, 0, 5},
		{"Negative shared attribute pairs", 0, -5, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CalcLCOM(tc.nonSharedAttributePairs, tc.sharedAttributePairs)

			if got != tc.expected {
				t.Fatalf("%v failed: results don't match. CalcDMS(%v, %v) == %v, expected %v", tc.name, tc.nonSharedAttributePairs, tc.sharedAttributePairs, got, tc.expected)
			}
		})
	}
}
