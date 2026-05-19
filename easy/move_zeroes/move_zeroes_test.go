package movezeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	assert.Equal(t, []int{0}, moveZeroes([]int{0}))
	assert.Equal(t, []int{3, 12, 0, 0, 0}, moveZeroes([]int{0, 0, 0, 3, 12}))
	assert.Equal(t, []int{3, 12, 0, 0, 0}, moveZeroes([]int{0, 0, 3, 12, 0}))
	assert.Equal(t, []int{1, 3, 12, 0, 0, 0}, moveZeroes([]int{0, 1, 0, 3, 0, 12}))
}
