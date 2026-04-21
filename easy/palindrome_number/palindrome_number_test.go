package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome(0))
	assert.Equal(t, true, isPalindrome(121))
	assert.Equal(t, true, isPalindrome(24542))
	assert.Equal(t, false, isPalindrome(-121))
	assert.Equal(t, false, isPalindrome(1038))
}
