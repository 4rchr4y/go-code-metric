package pdq

import (
	"fmt"
	"testing"
)

func TestCalcLCOM96b(t *testing.T) {
	t.Run("valid: input 1", func(t *testing.T) {
		expected := 0.4444444444444444
		got, _ := CalcLCOM96b(3, 3, []int{3, 1, 1})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
	t.Run("valid: input 2", func(t *testing.T) {
		expected := 0.6666666666666666
		got, _ := CalcLCOM96b(3, 3, []int{1, 1, 1})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
	t.Run("valid: input 3", func(t *testing.T) {
		expected := 0.4444444444444444
		got, _ := CalcLCOM96b(3, 3, []int{2, 2, 1})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: input 4", func(t *testing.T) {
		expected := 0.4444444444444444
		got, _ := CalcLCOM96b(3, 3, []int{2, 1, 2})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: input 5", func(t *testing.T) {
		expected := 0.4444444444444444
		got, _ := CalcLCOM96b(3, 3, []int{1, 2, 2})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: all methods accesses attributes", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcLCOM96b(3, 3, []int{3, 3, 3})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: no method accesses attributes", func(t *testing.T) {
		expected := 1.0
		got, _ := CalcLCOM96b(3, 3, []int{0, 0, 0})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: low num of attributes and high num of methods", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcLCOM96b(1, 9, []int{9})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: no attributes", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcLCOM96b(0, 3, []int{})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: no methods", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcLCOM96b(3, 0, []int{})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("invalid: attribute count mismatch", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("mismatch in numbers: attributesNum: %v is not equal to methodsPerAttribute: %v", 2, 3)
		got, err := CalcLCOM96b(2, 3, []int{1, 2, 1})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: negative attributes", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: attributesNum %v, methodsNum %v; values must not be less than 0", -1, 3)
		got, err := CalcLCOM96b(-1, 3, []int{})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: negative methods", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: attributesNum %v, methodsNum %v; values must not be less than 0", 3, -1)
		got, err := CalcLCOM96b(3, -1, []int{})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: negative value in methods per attribute slice 1", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid number of methods for attribute at index 0: -1; should be at least 0")
		got, err := CalcLCOM96b(3, 3, []int{-1, 0, 0})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: negative value in methods per attribute slice 2", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid number of methods for attribute at index 2: -1; should be at least 0")
		got, err := CalcLCOM96b(3, 3, []int{0, 0, -1})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("invalid: invalid methods per attribute exceeds total", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("too many methods for attribute at index %d: %d; there are only %d methods", 0, 99, 3)
		got, err := CalcLCOM96b(3, 3, []int{99, 1, 1})

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})
}
