package pdq

import (
	"fmt"
	"math"
	"testing"
)

func TestInstability(t *testing.T) {
	t.Run("valid: positive numbers", func(t *testing.T) {
		expected := 0.3333333333333333
		got, _ := CalcInstability(5, 10)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: zero abstract", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcInstability(0, 20)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: zero concrete", func(t *testing.T) {
		expected := 1.0
		got, _ := CalcInstability(10, 0)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: both values are zero", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcInstability(0, 0)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: math.MaxInt32 values", func(t *testing.T) {
		expected := 0.5
		got, _ := CalcInstability(math.MaxInt32, math.MaxInt32)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: abstract greater than concrete", func(t *testing.T) {
		expected := 0.75
		got, _ := CalcInstability(30, 10)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: concrete greater than concrete", func(t *testing.T) {
		expected := 0.25
		got, _ := CalcInstability(10, 30)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: equal numbers", func(t *testing.T) {
		expected := 0.5
		got, _ := CalcInstability(20, 20)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("invalid: negative abstract", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: outgoingDependenciesNum %v, incomingDependenciesNum %v; values must not be less than 0", -1, 1)
		got, err := CalcInstability(-1, 1)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: negative concrete", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: outgoingDependenciesNum %v, incomingDependenciesNum %v; values must not be less than 0", 1, -1)
		got, err := CalcInstability(1, -1)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: both numbers are negative", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: outgoingDependenciesNum %v, incomingDependenciesNum %v; values must not be less than 0", -1, -1)
		got, err := CalcInstability(-1, -1)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})
}
