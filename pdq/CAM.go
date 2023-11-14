package pdq

import (
	"fmt"
)

func CalcCAM(sharedPairs int, totalPossiblePairs int) (float64, error) {
	if sharedPairs < 0 || totalPossiblePairs < 0 {
		return 0, fmt.Errorf("shared pairs and total possible number of methods pairs cannot be negative, but got %d, %d", sharedPairs, totalPossiblePairs)
	}

	if sharedPairs == 0 || totalPossiblePairs == 0 {
		return 0, nil
	}

	return float64(sharedPairs) / float64(totalPossiblePairs), nil
}
