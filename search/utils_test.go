package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceContain(t *testing.T) {
	t.Run("TrueIfContains", func(t *testing.T) {
		assert.Equal(t, SliceContain([]string{"a", "b", "c"}, "b"), true)
	})

	t.Run("FalseIfNotContains", func(t *testing.T) {
		assert.Equal(t, SliceContain([]string{"a", "b", "c"}, "d"), false)
	})
}

func TestFindAnyExtension(t *testing.T) {
	t.Run("FalseIfSliceEmpty", func(t *testing.T) {
		assert.Equal(t, FindAnyExtension([]string{}), false)
	})

	t.Run("TrueIfAllExtensionsCharFirst", func(t *testing.T) {
		assert.Equal(t, FindAnyExtension([]string{"*"}), true)
	})
}
