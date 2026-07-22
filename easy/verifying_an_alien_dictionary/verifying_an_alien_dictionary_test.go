package verifyinganaliendictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlienSorted(t *testing.T) {
	assert.Equal(t, false, isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
	assert.Equal(t, true, isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	assert.Equal(t, false, isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
}
