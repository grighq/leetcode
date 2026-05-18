package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome(" "))
	assert.Equal(t, false, isPalindrome("0P"))
	assert.Equal(t, false, isPalindrome("race a car"))
	assert.Equal(t, true, isPalindrome("A man, a plan, a canal: Panama"))
}
