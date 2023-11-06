package pdq

import (
	"fmt"
)

func CalcLCOM96b(attributesNum int, methodsNum int, methodsPerAttribute []int) (float64, error) {
	if attributesNum < 0 {
		return -1, fmt.Errorf("invalid number of attributes: %d; It shouldn't be negative", attributesNum)
	}

	if methodsNum < 0 {
		return -1, fmt.Errorf("invalid number of methods: %d; It shouldn't be negative", methodsNum)
	}

	if attributesNum == 0 || methodsNum == 0 {
		return 1, nil
	}

	if attributesNum != len(methodsPerAttribute) {
		return -1, fmt.Errorf("mismatch in numbers: attributesNum is %d but methodsPerAttribute has %d elements; they should be equal", attributesNum, len(methodsPerAttribute))
	}

	methodsDifferenceSum := 0.0
	for i, methodsForAttribute := range methodsPerAttribute {
		if methodsForAttribute < 0 {
			return -1, fmt.Errorf("invalid number of methods for attribute at index %d: %d; should be at least 0", i+1, methodsForAttribute)
		}
		if methodsForAttribute > methodsNum {
			return -1, fmt.Errorf("invalid number of methods for attribute at index %d: %d; should be less or equal to the total number of methods: %d", i+1, methodsForAttribute, methodsNum)
		}

		methodsDifferenceSum += float64(methodsNum - methodsForAttribute)
	}

	lcom96b := (1.0 / float64(attributesNum)) * methodsDifferenceSum / float64(methodsNum)

	return lcom96b, nil
}
