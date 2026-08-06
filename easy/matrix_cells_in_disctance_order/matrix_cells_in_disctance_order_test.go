package matrixcellsindisctanceorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllCellsDistOrder(t *testing.T) {
	assert.Equal(t, [][]int{{0, 0}, {0, 1}}, allCellsDistOrder(1, 2, 0, 0))
	assert.Equal(t, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}, allCellsDistOrder(2, 2, 0, 1))
	assert.Equal(t, [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}}, allCellsDistOrder(2, 3, 1, 2))
}
