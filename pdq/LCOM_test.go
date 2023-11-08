package pdq

import (
	"math"
	"testing"
)

func TestCalcLCOM(t *testing.T) {
	t.Run("valid: more non-shared than shared attribute pairs", func(t *testing.T) {
		expected := 5
		got := CalcLCOM(10, 5)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: more shared than non-shared attribute pairs", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(5, 10)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
	t.Run("valid: equal shared and non-shared attribute pairs", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(5, 5)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: zero non-shared attribute pairs", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(0, 5)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: zero shared attribute pairs", func(t *testing.T) {
		expected := 5
		got := CalcLCOM(5, 0)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: both pairs are zero", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(0, 0)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: negative non-shared attribute pairs", func(t *testing.T) {
		expected := 5
		got := CalcLCOM(-5, 0)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: negative shared attribute pairs", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(0, -5)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: both pairs are math.MaxInt", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(math.MaxInt, math.MaxInt)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})

	t.Run("valid: both pairs are -math.MaxInt", func(t *testing.T) {
		expected := 0
		got := CalcLCOM(-math.MaxInt, -math.MaxInt)

		if got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}
