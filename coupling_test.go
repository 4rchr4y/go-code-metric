package gometric

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstability(t *testing.T) {
	t.Run("valid: positive numbers", func(t *testing.T) {
		got, err := CalcInstability(5, 10)

		assert.Equal(t, 0.3333333333333333, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero abstract", func(t *testing.T) {
		got, err := CalcInstability(0, 20)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: zero concrete", func(t *testing.T) {
		got, err := CalcInstability(10, 0)

		assert.Equal(t, 1.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: both values are zero", func(t *testing.T) {
		got, err := CalcInstability(0, 0)

		assert.Equal(t, 0.0, got)
		require.NoError(t, err)
	})

	t.Run("valid: math.MaxInt32 values", func(t *testing.T) {
		got, err := CalcInstability(math.MaxInt32, math.MaxInt32)

		assert.Equal(t, 0.5, got)
		require.NoError(t, err)
	})

	t.Run("valid: abstract greater than concrete", func(t *testing.T) {
		got, err := CalcInstability(30, 10)

		assert.Equal(t, 0.75, got)
		require.NoError(t, err)
	})

	t.Run("valid: concrete greater than concrete", func(t *testing.T) {
		got, err := CalcInstability(10, 30)

		assert.Equal(t, 0.25, got)
		require.NoError(t, err)
	})

	t.Run("valid: equal numbers", func(t *testing.T) {
		got, err := CalcInstability(20, 20)

		assert.Equal(t, 0.5, got)
		require.NoError(t, err)
	})

	t.Run("invalid: negative abstract", func(t *testing.T) {
		got, err := CalcInstability(-1, 1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative concrete", func(t *testing.T) {
		got, err := CalcInstability(1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: both numbers are negative", func(t *testing.T) {
		got, err := CalcInstability(-1, -1)

		assert.Equal(t, 0.0, got)
		require.Error(t, err)
	})
}
