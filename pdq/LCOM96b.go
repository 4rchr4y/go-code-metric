package pdq

import (
	"fmt"
)

func CalcLCOM96b(attributesNum int, methodsNum int, methodsPerAttribute []int) (float64, error) {
	if attributesNum < 0 || methodsNum < 0 {
		return -1, fmt.Errorf("attributesNum %v, methodsNum %v; values must be >= 0", attributesNum, methodsNum)
	}

	if attributesNum == 0 || methodsNum == 0 {
		return 0, nil
	}

	if attributesNum != len(methodsPerAttribute) {
		return -1, fmt.Errorf("attributesNum: %v != methodsPerAttribute length: %v", attributesNum, len(methodsPerAttribute))
	}

	methodsDifferenceSum := 0
	for i, methodsForAttribute := range methodsPerAttribute {
		if methodsForAttribute < 0 {
			return -1, fmt.Errorf("invalid number of methods for attribute at index %d: %d; should be >= 0", i, methodsForAttribute)
		}
		if methodsForAttribute > methodsNum {
			return -1, fmt.Errorf("too many methods for attribute at index %d: %d; should be <= methodsNum: %d", i, methodsForAttribute, methodsNum)
		}

		methodsDifferenceSum += methodsNum - methodsForAttribute
	}

	lcom96b := float64(methodsDifferenceSum) / (float64(attributesNum * methodsNum))

	return lcom96b, nil
}
