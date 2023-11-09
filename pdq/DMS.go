package pdq

import (
	"fmt"
	"math"
)

func CalcDMS(abstractness float64, instability float64) (float64, error) {
	if abstractness < 0 || instability < 0 {
		return -1, fmt.Errorf("abstractness and instability cannot be negative, but got %.2f, %.2f", abstractness, instability)
	}

	if abstractness >= math.Inf(1) || instability >= math.Inf(1) {
		return -1, fmt.Errorf("abstractness and instability must be lesser than infinity, but got %.2f, %.2f", abstractness, instability)
	}

	return math.Abs(abstractness + instability - 1), nil
}
