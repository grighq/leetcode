package removeduplicatesfromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	assert.Equal(t, 1, removeDuplicates([]int{3}))
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
