package sortarraybyparitiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParityII(t *testing.T) {
	assert.Equal(t, []int{2, 3}, sortArrayByParityII([]int{2, 3}))
	assert.Equal(t, []int{2, 9, 2, 3, 4, 3}, sortArrayByParityII([]int{2, 4, 2, 3, 9, 3}))
	assert.Equal(t, []int{2, 3, 0, 1, 4, 3}, sortArrayByParityII([]int{2, 0, 3, 4, 1, 3}))
	assert.Equal(t, []int{2, 3, 8, 3, 6, 3}, sortArrayByParityII([]int{3, 3, 3, 2, 6, 8}))
}
