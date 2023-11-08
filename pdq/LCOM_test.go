package pdq

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

		assert.Equal(t, -1, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative shared attribute pairs", func(t *testing.T) {
		got, err := CalcLCOM(0, -5)

		assert.Equal(t, -1, got)
		require.Error(t, err)
	})

	t.Run("invalid: both pairs are -math.MaxInt", func(t *testing.T) {
		got, err := CalcLCOM(-math.MaxInt, -math.MaxInt)

		assert.Equal(t, -1, got)
		require.Error(t, err)
	})
}
