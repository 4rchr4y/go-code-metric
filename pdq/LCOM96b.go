package pdq

import (
	"fmt"
)

func CalcLCOM96b(attributesNum int, methodsNum int, methodsPerAttribute []int) (float64, error) {
	if attributesNum < 0 || methodsNum < 0 {
		return -1, fmt.Errorf("invalid metric value: attributesNum %v, methodsNum %v; values must not be less than 0", attributesNum, methodsNum)
	}

	if attributesNum == 0 || methodsNum == 0 {
		return 0, nil
	}

	if attributesNum != len(methodsPerAttribute) {
		return -1, fmt.Errorf("mismatch in numbers: attributesNum: %v is not equal to methodsPerAttribute: %v", attributesNum, len(methodsPerAttribute))
	}

	methodsDifferenceSum := 0
	for i, methodsForAttribute := range methodsPerAttribute {
		if methodsForAttribute < 0 {
			return -1, fmt.Errorf("invalid number of methods for attribute at index %d: %d; should be at least 0", i, methodsForAttribute)
		}
		if methodsForAttribute > methodsNum {
			return -1, fmt.Errorf("too many methods for attribute at index %d: %d; there are only %d methods", i, methodsForAttribute, methodsNum)
		}

		methodsDifferenceSum += methodsNum - methodsForAttribute
	}

	lcom96b := float64(methodsDifferenceSum) / (float64(attributesNum * methodsNum))

	return lcom96b, nil
}
