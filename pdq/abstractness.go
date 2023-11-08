package pdq

import (
	"fmt"
)

func CalcAbstractness(abstractElementsNum int, concreteElementsNum int) (float64, error) {
	if abstractElementsNum < 0 || concreteElementsNum < 0 {
		return -1, fmt.Errorf("abstractElementsNum %v, concreteElementsNum %v; values must be >= 0", abstractElementsNum, concreteElementsNum)
	}

	if concreteElementsNum == 0 {
		return float64(abstractElementsNum), nil
	}

	abstractness := float64(abstractElementsNum) / float64(concreteElementsNum)

	return abstractness, nil
}
