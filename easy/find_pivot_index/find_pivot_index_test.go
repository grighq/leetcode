package findpivotindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPivotIndex(t *testing.T) {
	assert.Equal(t, 0, pivotIndex([]int{2, 1, -1}))
	assert.Equal(t, -1, pivotIndex([]int{1, 2, 3}))
	assert.Equal(t, 3, pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
