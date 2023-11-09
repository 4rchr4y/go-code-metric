package pdq

import (
	"fmt"
)

func CalcAbstractness(abstractElementsNum int, concreteElementsNum int) (float64, error) {
	if abstractElementsNum < 0 || concreteElementsNum < 0 {
		return 0, fmt.Errorf("number of abstract elements and concrete elements cannot be negative, but got %d, %d", abstractElementsNum, concreteElementsNum)
	}

	if concreteElementsNum == 0 {
		return float64(abstractElementsNum), nil
	}

	abstractness := float64(abstractElementsNum) / float64(concreteElementsNum)

	return abstractness, nil
}
