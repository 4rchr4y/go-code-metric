package gometric

import (
	"fmt"
)

// CalcAbstractness calculates the "abstractness" of a struct.
//
// Returns an error for negative input values.
//
// https://github.com/4rchr4y/go-code-metric#abstractness
func CalcAbstractness(abstractElementsNum int, concreteElementsNum int) (float64, error) {
	if abstractElementsNum < 0 || concreteElementsNum < 0 {
		return 0, fmt.Errorf("number of abstract elements and concrete elements cannot be negative, but got %d, %d", abstractElementsNum, concreteElementsNum)
	}

	totalElements := abstractElementsNum + concreteElementsNum
	if totalElements == 0 {
		return 0, nil
	}

	if concreteElementsNum == 0 {
		return 1, nil
	}

	abstractness := float64(abstractElementsNum) / float64(concreteElementsNum)

	return abstractness, nil
}
