package partitionarrayintothreepartswithequalsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanThreePartsEqualSum(t *testing.T) {
	assert.Equal(t, true, canThreePartsEqualSum([]int{1, -1, 1, -1, 1, -1, 1, -1}))
	assert.Equal(t, true, canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
	assert.Equal(t, true, canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	assert.Equal(t, false, canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
}
