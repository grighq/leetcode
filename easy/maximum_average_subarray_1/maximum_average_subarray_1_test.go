package maximumaveragesubarray1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumAverageSubarray1(t *testing.T) {
	assert.Equal(t, 5.0, findMaxAverage([]int{5}, 1))
	assert.Equal(t, 3.0, findMaxAverage([]int{4, 2, 1, 3, 3}, 2))
	assert.Equal(t, 4.0, findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
	assert.Equal(t, 12.75, findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
