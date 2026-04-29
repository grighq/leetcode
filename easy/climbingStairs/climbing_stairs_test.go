package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{6, 13},
	{45, 1836311903},
}

func TestClimbStairs(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, climbStairs(c.input))
	}
}

func TestClimbStairsRecursion(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, climbStairsRecursion(c.input))
	}
}
