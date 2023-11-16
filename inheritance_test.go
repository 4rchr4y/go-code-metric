package gometric

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAbstractness(t *testing.T) {
	t.Run("valid: input 1", func(t *testing.T) {
		got, err := CalcAbstractness(30, 10)

		assert.Equal(t, 3.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 2", func(t *testing.T) {
		got, err := CalcAbstractness(10, 30)

		assert.Equal(t, 0.3333333333333333, got)
		require.NoError(t, err)
	})

	t.Run("valid: input 3", func(t *testing.T) {
		got, err := CalcAbstractness(20, 20)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero abstract elements", func(t *testing.T) {
		got, err := CalcAbstractness(0, 20)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero concrete elements", func(t *testing.T) {
		got, err := CalcAbstractness(10, 0)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both values are math.MaxInt32", func(t *testing.T) {
		got, err := CalcAbstractness(math.MaxInt32, math.MaxInt32)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("invalid: input -1, 1", func(t *testing.T) {
		got, err := CalcAbstractness(-1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: input 1, -1,", func(t *testing.T) {
		got, err := CalcAbstractness(1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}
