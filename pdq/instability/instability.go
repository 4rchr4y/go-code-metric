package instability

import (
	"errors"
)

func CalcInstability(outgoingDependenciesNum int, incomingDependenciesNum int) (float64, error) {
	if outgoingDependenciesNum < 0 {
		return -1, errors.New("the number of outgoing dependencies must not be a negative number")
	}

	if incomingDependenciesNum < 0 {
		return -1, errors.New("the number of incoming dependencies must not be a negative number")
	}

	if outgoingDependenciesNum == 0 || incomingDependenciesNum == 0 {
		return 0, nil
	}

	instability := float64(outgoingDependenciesNum) / (float64(outgoingDependenciesNum) + float64(incomingDependenciesNum))

	return instability, nil
}
