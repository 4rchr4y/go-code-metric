package gometric

import (
	"fmt"
	"math"
)

// CalcLCOM calculates the Lack of Cohesion in Methods of a struct.
//
// Returns an error for negative input values.
//
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

// CalcLCOM96b calculates the Lack of Cohesion in Methods of a struct.
//
// Returns an error for negative input values
// or if attributes don't match the length of methods per attribute array,
// or if methods per attribute contain invalid values.
//
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

// CalcCAM calculates the Class Attribute Method of a struct.
//
// Returns an error for negative input values
// or if the total possible pairs is 0 but the number of shared pairs != 0.
//
// https://github.com/4rchr4y/go-code-metric#cohesion-among-methods-in-class-cam
func CalcCAM(sharedPairs int, totalPossiblePairs int) (float64, error) {
	if sharedPairs < 0 || totalPossiblePairs < 0 {
		return 0, fmt.Errorf("shared pairs and total possible number of methods pairs cannot be negative, but got %d, %d", sharedPairs, totalPossiblePairs)
	}

	if totalPossiblePairs == 0 {
		if sharedPairs == 0 {
			return 0, nil
		}

		return 0, fmt.Errorf("total possible number of methods pairs is zero, but, the number of method pairs is non-zero %d", sharedPairs)
	}

	return float64(sharedPairs) / float64(totalPossiblePairs), nil
}

// CalcReusabilityIndex calculates the reusability index of a struct.
//
// Returns an error for NaN input values
// or if cohesion, coupling, testability or documentation values are negative.
//
// https://github.com/4rchr4y/go-code-metric#reusability-index
func CalcReusabilityIndex(cohesion, coupling, testability, documentation, w1, w2, w3, w4 float64) (float64, error) {
	if math.IsNaN(cohesion) || math.IsNaN(coupling) || math.IsNaN(testability) || math.IsNaN(documentation) || math.IsNaN(w1) || math.IsNaN(w2) || math.IsNaN(w3) || math.IsNaN(w4) {
		return 0, fmt.Errorf("cohesion, coupling, testability, documentation or weights must not be NaN, but got: %.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f,%.2f", cohesion, coupling, testability, documentation, w1, w2, w3, w4)
	}

	if cohesion < 0 || coupling < 0 || testability < 0 || documentation < 0 {
		return 0, fmt.Errorf("cohesion, coupling, testability or documentation cannot be less than 0, but got: %.2f,%.2f,%.2f,%.2f", cohesion, coupling, testability, documentation)
	}

	return w1*cohesion + w2*(1-coupling) + w3*testability + w4*documentation, nil
}
