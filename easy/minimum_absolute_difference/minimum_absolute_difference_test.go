package minimumabsolutedifference

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumAbsDifference(t *testing.T) {
	assert.Equal(t, [][]int{[]int{1, 3}}, minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	assert.Equal(t, [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}}, minimumAbsDifference([]int{4, 2, 3, 1}))
	assert.Equal(t, [][]int{[]int{-14, -10}, []int{19, 23}, []int{23, 27}}, minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
