package pdq

import (
	"errors"
	"testing"
)

func TestCalcLCOM96b(t *testing.T) {
	testCases := []struct {
		name                string
		attributesNum       int
		methodsNum          int
		methodsPerAttribute []int
		expected            float64
		expectedError       error
	}{
		{"ValidInput1", 3, 3, []int{3, 1, 1}, 0.4444444444444444, nil},
		{"ValidInput2", 3, 3, []int{1, 1, 1}, 0.6666666666666666, nil},
		{"ValidInput3", 3, 3, []int{2, 2, 1}, 0.4444444444444444, nil},
		{"ValidInput4", 3, 3, []int{2, 1, 2}, 0.4444444444444444, nil},
		{"ValidInput5", 3, 3, []int{1, 2, 2}, 0.4444444444444444, nil},

		{"InvalidMethodsPerAttributeExceedsTotal", 3, 3, []int{99, 1, 1}, -1, errors.New("invalid number of methods for attribute at index 1: 99; should be less or equal to the total number of methods: 3")},

		{"AllMethodsAccessesAttributes", 3, 3, []int{3, 3, 3}, 0, nil},
		{"NoMethodAccessesAttributes", 3, 3, []int{0, 0, 0}, 1, nil},

		{"LowNumOfAttributesAndHighNumOfMethods", 1, 9, []int{9}, 0, nil},

		{"AttributeCountMismatch", 2, 3, []int{1, 2, 1}, -1, errors.New("mismatch in numbers: attributesNum is 2 but methodsPerAttribute has 3 elements; they should be equal")},

		{"NoAttributes", 0, 3, []int{}, 1, nil},
		{"NoMethods", 3, 0, []int{}, 1, nil},

		{"NegativeAttributes", -1, 3, []int{}, -1, errors.New("invalid number of attributes: -1; It shouldn't be negative")},
		{"NegativeMethods", 3, -1, []int{}, -1, errors.New("invalid number of methods: -1; It shouldn't be negative")},

		{"NegativeValueInMethodsPerAttribute1", 3, 3, []int{-1, 0, 0}, -1, errors.New("invalid number of methods for attribute at index 1: -1; should be at least 0")},
		{"NegativeValueInMethodsPerAttribute2", 3, 3, []int{0, 0, -1}, -1, errors.New("invalid number of methods for attribute at index 3: -1; should be at least 0")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalcLCOM96b(tc.attributesNum, tc.methodsNum, tc.methodsPerAttribute)

			if got != tc.expected {
				t.Fatalf("%v failed: results don't match. CalcLCOM96b(%v, %v, %v) == %v, expected %v", tc.name, tc.attributesNum, tc.methodsNum, tc.methodsPerAttribute, got, tc.expected)
			}

			if err != nil && err.Error() != tc.expectedError.Error() {
				t.Fatalf("%v failed: errors don't match. Error: %v. Expected error: %v", tc.name, err, tc.expectedError)
			}
		})
	}
}
