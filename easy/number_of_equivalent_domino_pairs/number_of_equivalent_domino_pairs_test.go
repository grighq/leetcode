package numberofequivalentdominopairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumEquivDominoPairs(t *testing.T) {
	assert.Equal(t, 1, numEquivDominoPairs([][]int{[]int{1, 2}, []int{2, 1}, []int{3, 4}, []int{5, 6}}))
	assert.Equal(t, 3, numEquivDominoPairs([][]int{[]int{1, 2}, []int{1, 2}, []int{1, 1}, []int{1, 2}, []int{2, 2}}))
}
