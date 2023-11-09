package pdq

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDMS(t *testing.T) {
	t.Run("valid: values at lower bound", func(t *testing.T) {
		got, err := CalcDMS(0, 0)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: values at upper bound", func(t *testing.T) {
		got, err := CalcDMS(1, 1)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: equal values", func(t *testing.T) {
		got, err := CalcDMS(0.5, 0.5)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: complimentary values", func(t *testing.T) {
		got, err := CalcDMS(0.25, 0.75)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: arbitrary values", func(t *testing.T) {
		got, err := CalcDMS(0.1, 0.2)

		assert.Equal(t, 0.7, got)
		require.NoError(t, err)
	})

	t.Run("valid: near-equal values", func(t *testing.T) {
		got, err := CalcDMS(0.50000000000001, 0.49999999999999)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("Invalid: negative abstractness", func(t *testing.T) {
		got, err := CalcDMS(-1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("Invalid: negative instability", func(t *testing.T) {
		got, err := CalcDMS(1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("Invalid: abstractness is NaN", func(t *testing.T) {
		got, err := CalcDMS(math.NaN(), 0)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("Invalid: instability is NaN", func(t *testing.T) {
		got, err := CalcDMS(0, math.NaN())

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}
