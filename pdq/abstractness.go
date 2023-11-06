package pdq

import (
	"errors"
)

func CalcAbstractness(abstractElementsNum int, concreteElementsNum int) (float64, error) {
	if abstractElementsNum < 0 {
		return -1, errors.New("the number of abstract elements must not be a negative number")
	}

	if concreteElementsNum < 0 {
		return -1, errors.New("the number of concrete elements must not be a negative number")
	}

	if abstractElementsNum == 0 || concreteElementsNum == 0 {
		return 0, nil
	}

	abstractness := float64(abstractElementsNum) / (float64(abstractElementsNum) + float64(concreteElementsNum))

	return abstractness, nil
}
