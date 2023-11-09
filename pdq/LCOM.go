package pdq

import (
	"fmt"
)

func CalcLCOM(nonSharedFields int, sharedFields int) (int, error) {
	if nonSharedFields < 0 || sharedFields < 0 {
		return -1, fmt.Errorf("nonSharedFields %v, nonSharedFields %v; values must be >= 0", nonSharedFields, sharedFields)
	}

	if nonSharedFields > sharedFields {
		return nonSharedFields - sharedFields, nil
	}

	return 0, nil
}

func CalcLCOM96b(attributesNum int, methodsNum int, methodsPerAttribute []int) (float64, error) {
	if attributesNum < 0 || methodsNum < 0 {
		return -1, fmt.Errorf("attributesNum %d, methodsNum %d; values must be >= 0", attributesNum, methodsNum)
	}

	if attributesNum == 0 || methodsNum == 0 {
		return 0, nil
	}

	if attributesNum != len(methodsPerAttribute) {
		return -1, fmt.Errorf("attributesNum: %d != methodsPerAttribute length: %d", attributesNum, len(methodsPerAttribute))
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
