package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathExist(t *testing.T) {
	t.Parallel()

	t.Run("want got true", func(t *testing.T) {

		expected := PathExist(".")

		assert.Equal(t, expected, true)

	})

	t.Run("want got false", func(t *testing.T) {

		expected := PathExist("lhhajja/asasas")

		assert.Equal(t, expected, false)

	})
}

func TestAlphaNum(t *testing.T) {
	t.Run("Test when right alpha num", func(t *testing.T) {

		e := AlphaNum("acbxsddd1212313")
		assert.NoError(t, e)
	})

	t.Run("Test when empty string", func(t *testing.T) {

		e := AlphaNum("")
		assert.NoError(t, e)
	})

	t.Run("Test when wrong alpha num", func(t *testing.T) {
		e := AlphaNum("acbxsddd1212313--566    $%%$##\"")
		assert.Error(t, e)
	})
}
