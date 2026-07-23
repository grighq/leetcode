package largestperimetertriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestPerimeter(t *testing.T) {
	assert.Equal(t, 5, largestPerimeter([]int{2, 1, 2}))
	assert.Equal(t, 0, largestPerimeter([]int{1, 2, 1, 10}))
}
