package DMS

import (
	"errors"
	"math"
)

func CalcDMS(abstractness float64, instability float64) (float64, error) {
	if abstractness < 0 {
		return -1, errors.New("abstractness cannot be lesser than 0")
	}

	if instability < 0 {
		return -1, errors.New("instability cannot be lesser than 0")
	}

	return math.Abs(abstractness + instability - 1), nil
}
