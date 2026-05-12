package fibonaccinumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
}

func TestFib(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, fib(c.input))
	}
}
func TestFibRecursion(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, fibRecurcsion(c.input))
	}
}
