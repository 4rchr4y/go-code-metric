package pdq

import (
	"fmt"
	"math"
)

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
