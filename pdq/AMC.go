package pdq

import (
	"fmt"
)

func CalcAMC(methodComplexity int, totalMethodsNum int) (float64, error) {
	if methodComplexity < 0 || totalMethodsNum < 0 {
		return 0, fmt.Errorf("method complexity and total number of methods cannot be negative, but got %d, %d", methodComplexity, totalMethodsNum)
	}

	if methodComplexity == 0 || totalMethodsNum == 0 {
		return 0, nil
	}

	return float64(methodComplexity) / float64(totalMethodsNum), nil
}
