package pdq

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

		assert.Equal(t, -1.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative attributes", func(t *testing.T) {
		got, err := CalcLCOM96b(-1, 3, []int{})

		assert.Equal(t, -1.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative methods", func(t *testing.T) {
		got, err := CalcLCOM96b(3, -1, []int{})

		assert.Equal(t, -1.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative value in methods per attribute slice 1", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{-1, 0, 0})

		assert.Equal(t, -1.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: negative value in methods per attribute slice 2", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{0, 0, -1})

		assert.Equal(t, -1.0, got)
		require.Error(t, err)
	})

	t.Run("invalid: invalid methods per attribute exceeds total", func(t *testing.T) {
		got, err := CalcLCOM96b(3, 3, []int{99, 1, 1})

		assert.Equal(t, -1.0, got)
		require.Error(t, err)
	})
}
