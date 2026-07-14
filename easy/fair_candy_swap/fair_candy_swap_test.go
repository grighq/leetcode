package faircandyswap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFairCandySwap(t *testing.T) {
	assert.Equal(t, []int{2, 3}, fairCandySwap([]int{2}, []int{1, 3}))
	assert.Equal(t, []int{1, 2}, fairCandySwap([]int{1, 1}, []int{2, 2}))
	assert.Equal(t, []int{1, 2}, fairCandySwap([]int{1, 1}, []int{2, 3}))
	assert.Equal(t, []int{5, 4}, fairCandySwap([]int{1, 2, 5}, []int{2, 4}))
}
