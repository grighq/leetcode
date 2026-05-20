package reversestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    []byte
	expected []byte
}{
	{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
}

func TestReverseString(t *testing.T) {
	for _, c := range cases {
		s := c.input
		reverseString(s)
		assert.Equal(t, c.expected, s)
	}
}
