package instability

import (
	"errors"
	"math"
	"testing"
)

func TestInstability(t *testing.T) {
	tests := []struct {
		name                    string
		outgoingDependenciesNum int
		incomingDependenciesNum int
		expected                float64
		expectedError           error
	}{
		{"PositiveNumbers", 10, 20, 0.3333333333333333, nil},
		{"ZeroAbstractNumber", 0, 20, 0, nil},
		{"ZeroConcreteNumber", 10, 0, 0, nil},
		{"NegativeAbstractNumber", -1, 1, -1, errors.New("the number of outgoing dependencies must not be a negative number")},
		{"NegativeConcreteNumber", 1, -1, -1, errors.New("the number of incoming dependencies must not be a negative number")},
		{"BothNumbersNegative", -10, -20, -1, errors.New("the number of outgoing dependencies must not be a negative number")},
		{"LargeNumbers", math.MaxInt32, math.MaxInt32, 0.5, nil},
		{"AbstractGreaterThanConcrete", 30, 10, 0.75, nil},
		{"ConcreteGreaterThanAbstract", 10, 30, 0.25, nil},
		{"EqualNumbers", 20, 20, 0.5, nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalcInstability(tc.outgoingDependenciesNum, tc.incomingDependenciesNum)

			if got != tc.expected {
				t.Fatalf("%v failed: results don't match. CalcInstability(%v, %v)== %v. Expected: %v", tc.name, tc.outgoingDependenciesNum, tc.incomingDependenciesNum, got, tc.expected)
			}

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Fatalf("%v failed: errors don't match. Error: %v. Expected error: %v", tc.name, err, tc.expectedError)
			}
		})
	}
}
