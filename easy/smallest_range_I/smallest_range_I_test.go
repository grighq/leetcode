package smallestrangeI

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRangeI(t *testing.T) {
	assert.Equal(t, 0, smallestRangeI([]int{1}, 0))
	assert.Equal(t, 6, smallestRangeI([]int{0, 10}, 2))
	assert.Equal(t, 0, smallestRangeI([]int{1, 3, 6}, 3))
}
