package reversevowelsofastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowelsOfAString(t *testing.T) {
	assert.Equal(t, "AceCreIm", reverseVowels("IceCreAm"))
	assert.Equal(t, "leotcede", reverseVowels("leetcode"))
	assert.Equal(t, "Yo! Bottoms Up, u.S. Motto, boy!", reverseVowels("Yo! Bottoms up, U.S. Motto, boy!"))
}
