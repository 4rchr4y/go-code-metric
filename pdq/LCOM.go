package pdq

import "math"

func CalcLCOM(nonSharedAttributePairs int, sharedAttributePairs int) int {
	absoluteNonSharedPairs := math.Abs(float64(nonSharedAttributePairs))
	absoluteSharedPair := math.Abs(float64(sharedAttributePairs))

	if absoluteNonSharedPairs >= absoluteSharedPair {
		return int(absoluteNonSharedPairs - absoluteSharedPair)
	}

	return 0
}
