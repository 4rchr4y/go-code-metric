package pdq

import (
	"fmt"
)

func CalcInstability(outgoingDependenciesNum int, incomingDependenciesNum int) (float64, error) {
	if outgoingDependenciesNum < 0 || incomingDependenciesNum < 0 {
		return -1, fmt.Errorf("number of outgoing and incoming dependencies cannot be negative, but got %d, %d", outgoingDependenciesNum, incomingDependenciesNum)
	}

	totalDependencies := outgoingDependenciesNum + incomingDependenciesNum
	if totalDependencies == 0 {
		return 0, nil
	}

	instability := float64(outgoingDependenciesNum) / float64(totalDependencies)

	return instability, nil
}
