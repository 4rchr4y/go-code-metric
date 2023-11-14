package pdq

import (
	"fmt"
)

// https://github.com/4rchr4y/go-code-metric#average-method-complexity-amc

func CalcAMC(methodsComplexity int, totalMethodsNum int) (float64, error) {
	if methodsComplexity < 0 || totalMethodsNum < 0 {
		return 0, fmt.Errorf("method complexity and total number of methods cannot be negative, but got %d, %d", methodsComplexity, totalMethodsNum)
	}

	if methodsComplexity == 0 || totalMethodsNum == 0 {
		return 0, nil
	}

	return float64(methodsComplexity) / float64(totalMethodsNum), nil
}
