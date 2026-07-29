package maximizesumofarrayafterknegations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	assert.Equal(t, 5, largestSumAfterKNegations([]int{4, 2, 3}, 1))
	assert.Equal(t, 6, largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	assert.Equal(t, 13, largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}
