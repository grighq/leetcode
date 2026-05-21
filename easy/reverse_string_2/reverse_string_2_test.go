package reversestring2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStr(t *testing.T) {
	assert.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	assert.Equal(t, "bacd", reverseStr("abcd", 2))
}
