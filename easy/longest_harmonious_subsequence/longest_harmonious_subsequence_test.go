package longestharmonioussubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLHS(t *testing.T) {
	assert.Equal(t, 0, findLHS([]int{1, 1, 1, 1}))
	assert.Equal(t, 2, findLHS([]int{1, 2, 3, 4}))
	assert.Equal(t, 5, findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
