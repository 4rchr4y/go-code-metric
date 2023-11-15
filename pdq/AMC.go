package pdq

import (
	"fmt"
)

// CalcAMC calculates the Average Method Complexity of a struct.
//
// Returns an error for negative input values
// or if the total number of methods is 0 but the methods complexity != 0.
//
// https://github.com/4rchr4y/go-code-metric#average-method-complexity-amc
func CalcAMC(methodsComplexity int, totalMethodsNum int) (float64, error) {
	if methodsComplexity < 0 || totalMethodsNum < 0 {
		return 0, fmt.Errorf("method complexity and total number of methods cannot be negative, but got %d, %d", methodsComplexity, totalMethodsNum)
	}

	if totalMethodsNum == 0 {
		if methodsComplexity == 0 {
			return 0, nil
		}

		return 0, fmt.Errorf("total number of methods is 0, but, the methods complexity is non-zero %d", methodsComplexity)
	}

	return float64(methodsComplexity) / float64(totalMethodsNum), nil
}
