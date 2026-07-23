package addtoarrayformofinteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToArrayForm(t *testing.T) {
	assert.Equal(t, []int{1, 0, 0}, addToArrayForm([]int{1}, 99))
	assert.Equal(t, []int{4, 5, 5}, addToArrayForm([]int{2, 7, 4}, 181))
	assert.Equal(t, []int{1, 0, 2, 1}, addToArrayForm([]int{2, 1, 5}, 806))
	assert.Equal(t, []int{1, 2, 3, 4}, addToArrayForm([]int{1, 2, 0, 0}, 34))
}
