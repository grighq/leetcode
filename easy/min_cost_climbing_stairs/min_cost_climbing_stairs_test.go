package mincostclimbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	assert.Equal(t, 10, minCostClimbingStairs([]int{10, 15}))
	assert.Equal(t, 15, minCostClimbingStairs([]int{10, 15, 20}))
	assert.Equal(t, 6, minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
