package sortedarraybyparity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayByParity(t *testing.T) {
	assert.Equal(t, []int{0}, sortArrayByParity([]int{0}))
	assert.Equal(t, []int{2, 4, 3, 1}, sortArrayByParity([]int{3, 1, 2, 4}))
	assert.Equal(t, []int{4, 6, 3, 1}, sortArrayByParity([]int{4, 1, 3, 6}))
}
