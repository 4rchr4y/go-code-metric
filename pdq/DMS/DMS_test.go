package DMS

import (
	"errors"
	"math"
	"testing"
)

func TestDMS(t *testing.T) {
	tests := []struct {
		name          string
		abstractness  float64
		instability   float64
		expected      float64
		expectedError error
	}{
		{"Values at lower bound", 0, 0, 1, nil},
		{"Values at upper bound", 1, 1, 1, nil},
		{"Both values at equilibrium", 0.5, 0.5, 0, nil},
		{"Complimentary values", 0.25, 0.75, 0, nil},
		{"Arbitrary values", 0.1, 0.2, 0.7, nil},
		{"Near-equal values", 0.50000000000001, 0.49999999999999, 0, nil},
		{"Extreme large values", 1e308, 1e308, math.Inf(1), nil},
		{"Negative abstractness", -1, 1, -1, errors.New("abstractness cannot be lesser than 0")},
		{"Negative instability", 1, -1, -1, errors.New("instability cannot be lesser than 0")},
		{"Extreme small values", -1e308, -1e308, -1, errors.New("abstractness cannot be lesser than 0")},
		{"Abstractness infinity positive", math.Inf(1), 0, -1, errors.New("abstractness is infinite")},
		{"Abstractness infinity negative", math.Inf(-1), 0, -1, errors.New("abstractness is infinite")},
		{"Instability infinity positive", 0, math.Inf(1), -1, errors.New("instability is infinite")},
		{"Instability infinity negative", 0, math.Inf(-1), -1, errors.New("instability is infinite")},
		{"Abstractness NaN", math.NaN(), 0.2, -1, errors.New("abstractness is NaN")},
		{"Instability NaN", 0.2, math.NaN(), -1, errors.New("instability is NaN")},
		{"Abstractness and instability are NaN", math.NaN(), math.NaN(), -1, errors.New("abstractness is NaN")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalcDMS(tc.abstractness, tc.instability)

			if got != tc.expected {
				t.Fatalf("%v failed: results don't match. CalcDMS(%v, %v) == %v, expected %v", tc.name, tc.abstractness, tc.instability, got, tc.expected)
			}

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Fatalf("%v failed: errors don't match. Error: %v. Expected error: %v", tc.name, err, tc.expectedError)
			}
		})
	}
}
