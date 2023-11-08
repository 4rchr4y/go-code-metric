package pdq

import (
	"fmt"
	"math"
)

func CalcDMS(abstractness float64, instability float64) (float64, error) {
	if abstractness < 0 || instability < 0 {
		return -1, fmt.Errorf("abstractness %v, instability %v; values must be >= 0", abstractness, instability)
	}

	if abstractness >= 1e308 || instability >= 1e308 {
		return -1, fmt.Errorf("abstractness %v, instability %v; values must be < infinity", abstractness, instability)
	}

	return math.Abs(abstractness + instability - 1), nil
}
