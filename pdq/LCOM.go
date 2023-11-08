package pdq

import (
	"fmt"
)

func CalcLCOM(nonSharedFields int, sharedFields int) (int, error) {
	if nonSharedFields < 0 || sharedFields < 0 {
		return -1, fmt.Errorf("nonSharedFields %v, nonSharedFields %v; values must be >= 0", nonSharedFields, sharedFields)
	}

	if nonSharedFields > sharedFields {
		return nonSharedFields - sharedFields, nil
	}

	return 0, nil
}
