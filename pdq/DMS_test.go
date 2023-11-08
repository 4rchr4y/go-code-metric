package pdq

import (
	"fmt"
	"testing"
)

func TestDMS(t *testing.T) {
	t.Run("Valid: values at lower bound", func(t *testing.T) {
		expected := 1.0
		got, _ := CalcDMS(0, 0)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Valid: values at upper bound", func(t *testing.T) {
		expected := 1.0
		got, _ := CalcDMS(1, 1)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Valid: equal values", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcDMS(0.5, 0.5)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Valid: complimentary values", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcDMS(0.25, 0.75)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Valid: arbitrary values", func(t *testing.T) {
		expected := 0.7
		got, _ := CalcDMS(0.1, 0.2)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Valid: near-equal values", func(t *testing.T) {
		expected := 0.0
		got, _ := CalcDMS(0.50000000000001, 0.49999999999999)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("Invalid: negative abstractness", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractness %v, instability %v; values must be 0 or higher", -1, 1)
		got, err := CalcDMS(-1, 1)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("Invalid: negative instability", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractness %v, instability %v; values must be 0 or higher", 1, -1)
		got, err := CalcDMS(1, -1)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("Invalid: abstractness is 1e308", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractness %v, instability %v; values must be lesser than infinity", 1e308, 0)
		got, err := CalcDMS(1e308, 0)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("Invalid: instability is 1e308", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractness %v, instability %v; values must be lesser than infinity", 0, 1e308)
		got, err := CalcDMS(0, 1e308)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("Invalid: abstractness is -1e308", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractness %v, instability %v; values must be 0 or higher", -1e308, 0)
		got, err := CalcDMS(-1e308, 0)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})

	t.Run("Invalid: instability is -1e308", func(t *testing.T) {
		expected := -1.0
		expectedErr := fmt.Errorf("invalid metric value: abstractness %v, instability %v; values must be 0 or higher", 0, -1e308)
		got, err := CalcDMS(0, -1e308)

		if got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}

		if err.Error() != expectedErr.Error() {
			t.Errorf("errors don't match: error (%v), expectedErr (%v)", err, expectedErr)
		}
	})
}
