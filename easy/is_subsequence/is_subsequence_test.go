package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	assert.Equal(t, false, isSubsequence("c", ""))
	assert.Equal(t, true, isSubsequence("", "hcd"))
	assert.Equal(t, true, isSubsequence("c", "hcd"))
	assert.Equal(t, true, isSubsequence("abc", "ahbgcd"))
	assert.Equal(t, false, isSubsequence("axc", "ahbgcd"))
}
