package gometric

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcLCOM(t *testing.T) {
	t.Run("valid: more non-shared than shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(10, 5)

		assert.Equal(t, 5, got)
		require.NoError(t, err)
	})

	t.Run("valid: more shared than non-shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(5, 10)

		assert.Equal(t, 0, got)
		require.NoError(t, err)
	})
	t.Run("valid: equal shared and non-shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(5, 5)

		assert.Equal(t, 0, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero non-shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(0, 5)

		assert.Equal(t, 0, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(5, 0)

		assert.Equal(t, 5, got)
		require.NoError(t, err)
	})

	t.Run("valid: both pairs are zero", func(t *testing.T) {
		got, err := CalcLCOM(0, 0)

		assert.Equal(t, 0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both pairs are math.MaxInt", func(t *testing.T) {
		got, err := CalcLCOM(math.MaxInt, math.MaxInt)

		assert.Equal(t, 0, got)
		require.NoError(t, err)
	})

	t.Run("invalid: negative non-shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(-5, 0)

		assert.Equal(t, 0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(0, -5)

		assert.Equal(t, 0, got)
		require.Error(t, err)
	})

	t.Run("invalid: both pairs are -math.MaxInt", func(t *testing.T) {
		got, err := CalcLCOM(-math.MaxInt, -math.MaxInt)

		assert.Equal(t, 0, got)
		require.Error(t, err)
	})
}

func TestCalcLCOM96b(t *testing.T) {
	t.Run("valid: input 1", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{3, 1, 1})

		assert.Equal(t, 0.4444444444444444, got)
		require.NoError(t, err)
	})
	t.Run("valid: input 2", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{1, 1, 1})

		assert.Equal(t, 0.6666666666666666, got)
		require.NoError(t, err)
	})
	t.Run("valid: input 3", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{2, 2, 1})

		assert.Equal(t, 0.4444444444444444, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 4", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{2, 1, 2})

		assert.Equal(t, 0.4444444444444444, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 5", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{1, 2, 2})

		assert.Equal(t, 0.4444444444444444, got)
		require.NoError(t, err)
	})

	t.Run("valid: all methods accesses attributes", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{3, 3, 3})

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: no method accesses attributes", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{0, 0, 0})

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: low num of attributes and high num of methods", func(t *testing.T) {
		got, err := CalcLCOM96b(1, 9, []int{9})

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: no attributes", func(t *testing.T) {
		got, err := CalcLCOM96b(0, 3, []int{})

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: no methods", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 0, []int{})

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("invalid: attribute count mismatch", func(t *testing.T) {
		got, err := CalcLCOM96b(2, 3, []int{1, 2, 1})

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative attributes", func(t *testing.T) {
		got, err := CalcLCOM96b(-1, 3, []int{})

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative methods", func(t *testing.T) {
		got, err := CalcLCOM96b(3, -1, []int{})

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative value in methods per attribute slice 1", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{-1, 0, 0})

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative value in methods per attribute slice 2", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{0, 0, -1})

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: invalid methods per attribute exceeds total", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{99, 1, 1})

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}

func TestCAM(t *testing.T) {
	t.Run("valid: input 1", func(t *testing.T) {
		got, err := CalcCAM(2, 2)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 2", func(t *testing.T) {
		got, err := CalcCAM(16, 5)

		assert.Equal(t, 3.2, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 3", func(t *testing.T) {
		got, err := CalcCAM(12, 3)

		assert.Equal(t, 4.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 4", func(t *testing.T) {
		got, err := CalcCAM(5, 3)

		assert.Equal(t, 1.6666666666666667, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 5", func(t *testing.T) {
		got, err := CalcCAM(17, 3)

		assert.Equal(t, 5.666666666666667, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero method complexity", func(t *testing.T) {
		got, err := CalcCAM(0, 20)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both values are 0", func(t *testing.T) {
		got, err := CalcCAM(0, 0)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both values are math.MaxInt32", func(t *testing.T) {
		got, err := CalcCAM(math.MaxInt32, math.MaxInt32)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("invalid: zero total number of methods", func(t *testing.T) {
		got, err := CalcCAM(10, 0)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative shared pairs", func(t *testing.T) {
		got, err := CalcCAM(-1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative total possible number of methods pairs", func(t *testing.T) {
		got, err := CalcCAM(1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}

func TestCalcReusabilityIndex(t *testing.T) {
	t.Run("valid: all inputs are 1", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, 3.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: all parameters are 1, all weights are 0.25", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, 1, 1, 0.25, 0.25, 0.25, 0.25)

		assert.Equal(t, 0.75, got)
		require.NoError(t, err)
	})

	t.Run("valid: all inputs are 0", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0, 0, 0, 0, 0, 0, 0, 0)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: expected result is 0", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0, 1, 0, 0, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: all parameters are 0.5, all weights are 0.25", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0.5, 0.5, 0.5, 0.5, 0.25, 0.25, 0.25, 0.25)

		assert.Equal(t, 0.5, got)
		require.NoError(t, err)
	})

	t.Run("valid: all parameters are 1, all weights are 0", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, 1, 1, 0, 0, 0, 0)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: random valid inputs 1", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0.9, 0.1, 0.8, 0.7, 0.2, 0.3, 0.4, 0.1)

		assert.Equal(t, 0.8400000000000001, got)
		require.NoError(t, err)
	})

	t.Run("valid: random valid inputs 2", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0.2, 0.8, 0.6, 0.4, 0.25, 0.25, 0.25, 0.25)

		assert.Equal(t, 0.35, got)
		require.NoError(t, err)
	})

	t.Run("valid: floating point precision", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0.333, 0.333, 0.333, 0.333, 0.25, 0.25, 0.25, 0.25)

		assert.Equal(t, 0.4165, got)
		require.NoError(t, err)
	})

	t.Run("valid: one weight significantly higher than others", func(t *testing.T) {
		got, err := CalcReusabilityIndex(0.5, 0.5, 0.5, 0.5, 0.1, 0.1, 0.1, 0.7)

		assert.Equal(t, 0.5, got)
		require.NoError(t, err)
	})

	t.Run("invalid: cohesion is negative", func(t *testing.T) {
		got, err := CalcReusabilityIndex(-1, 1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: coupling is negative", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, -1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: testability is negative", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, -1, 1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: documentation is negative", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, 1, -1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: cohesion is NaN", func(t *testing.T) {
		got, err := CalcReusabilityIndex(math.NaN(), 1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: coupling is NaN", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, math.NaN(), 1, 1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: testability is NaN", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, math.NaN(), 1, 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: documentation is NaN", func(t *testing.T) {
		got, err := CalcReusabilityIndex(1, 1, 1, math.NaN(), 1, 1, 1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}
