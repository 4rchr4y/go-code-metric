package pdq

import (
	"fmt"
	"math"
	"testing"
)

func TestAbstractness(t *testing.T) {

	t.Run("valid: input 1", func(t *testing.T) {
		expected := 3.0
		got, _ := CalcAbstractness(30, 10)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: input 2", func(t *testing.T) {
		expected := 0.3333333333333333
		got, _ := CalcAbstractness(10, 30)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: input 3", func(t *testing.T) {
		expected := 1.0
		got, _ := CalcAbstractness(20, 20)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: input 4", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcAbstractness(0, 20)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: input 5", func(t *testing.T) {
		expected := 10.0
		got, _ := CalcAbstractness(10, 0)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: both values are math.MaxInt32", func(t *testing.T) {
		expected := 1.0
		got, _ := CalcAbstractness(math.MaxInt32, math.MaxInt32)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Invalid: input -1, 1", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractElementsNum %v, concreteElementsNum %v; values must not be less than 0", -1, 1)
		got, err := CalcAbstractness(-1, 1)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("Invalid: input 1, -1,", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractElementsNum %v, concreteElementsNum %v; values must not be less than 0", 1, -1)
		got, err := CalcAbstractness(1, -1)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})
}
