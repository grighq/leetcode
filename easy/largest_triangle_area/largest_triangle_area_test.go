package largesttrianglearea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestTriangleArea(t *testing.T) {
	assert.Equal(t, 0.5, largestTriangleArea([][]int{{1, 0}, {0, 0}, {0, 1}}))
	assert.Equal(t, 2.0, largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}))
}
