package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsertPosition(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{1}, 0))
	assert.Equal(t, 1, searchInsert([]int{1}, 8))
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 4))
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
}
