package minimumcosttomovechipstothesameposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostToMoveChips(t *testing.T) {
	assert.Equal(t, 1, minCostToMoveChips([]int{1, 2, 3}))
	assert.Equal(t, 2, minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	assert.Equal(t, 1, minCostToMoveChips([]int{1, 1000000000}))
}
