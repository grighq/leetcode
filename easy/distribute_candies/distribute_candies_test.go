package distributecandies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	assert.Equal(t, 1, distributeCandies([]int{6, 6, 6, 6}))
	assert.Equal(t, 2, distributeCandies([]int{1, 1, 2, 3}))
	assert.Equal(t, 3, distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	assert.Equal(t, 2, distributeCandies([]int{100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0}))
}
