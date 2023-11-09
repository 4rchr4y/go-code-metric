package pdq

import (
	"fmt"
	"math"
)

func CalcDMS(abstractness float64, instability float64) (float64, error) {
	if math.IsNaN(abstractness) || math.IsNaN(instability) {
		return 0, fmt.Errorf("abstractness and instability must not be NaN, but got %.2f, %.2f", abstractness, instability)
	}

	if abstractness < 0 || instability < 0 {
		return 0, fmt.Errorf("abstractness and instability cannot be negative, but got %.2f, %.2f", abstractness, instability)
	}

	return math.Abs(abstractness + instability - 1), nil
}
