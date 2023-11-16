package gometric

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
