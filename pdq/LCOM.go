package pdq

import (
	"fmt"
)

// https://github.com/4rchr4y/go-code-metric#lack-of-cohesion-in-methodslcom

func CalcLCOM(nonSharedFields int, sharedFields int) (int, error) {
	if nonSharedFields < 0 || sharedFields < 0 {
		return 0, fmt.Errorf("nonSharedFields %v, nonSharedFields %v; values must be >= 0", nonSharedFields, sharedFields)
	}

	if nonSharedFields > sharedFields {
		return nonSharedFields - sharedFields, nil
	}

	return 0, nil
}

// https://github.com/4rchr4y/go-code-metric#lack-of-cohesion-in-methodslcom96b

func CalcLCOM96b(attributesNum int, methodsNum int, methodsPerAttribute []int) (float64, error) {
	if attributesNum < 0 || methodsNum < 0 {
		return 0, fmt.Errorf("number of attributes and methods cannot be negative, but got %d, %d", attributesNum, methodsNum)
	}

	if attributesNum == 0 || methodsNum == 0 {
		return 0, nil
	}

	if attributesNum != len(methodsPerAttribute) {
		return 0, fmt.Errorf("number of attributes should be equal to methods per attribute, but got %d, %d", attributesNum, len(methodsPerAttribute))
	}

	methodsDifferenceSum := 0
	for i, methodsForAttribute := range methodsPerAttribute {
		if methodsForAttribute < 0 {
			return 0, fmt.Errorf("number of methods for attribute at index %d: %d, should be >= 0", i, methodsForAttribute)
		}
		if methodsForAttribute > methodsNum {
			return 0, fmt.Errorf("too many methods for attribute at index %d: %d, should be <= than methodsNum: %d", i, methodsForAttribute, methodsNum)
		}

		methodsDifferenceSum += methodsNum - methodsForAttribute
	}

	lcom96b := float64(methodsDifferenceSum) / (float64(attributesNum * methodsNum))

	return lcom96b, nil
}
