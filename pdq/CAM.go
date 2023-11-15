package pdq

import (
	"fmt"
)

// CalcCAM calculates the Class Attribute Method of a struct.
// Returns an error for negative input values
// or if the total possible pairs is 0 but the number of shared pairs != 0.
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
