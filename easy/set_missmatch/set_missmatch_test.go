package setmissmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindErrorNums(t *testing.T) {
	assert.Equal(t, []int{1, 2}, findErrorNums([]int{1, 1}))
	assert.Equal(t, []int{2, 1}, findErrorNums([]int{2, 2}))
	assert.Equal(t, []int{2, 3}, findErrorNums([]int{1, 2, 2, 4}))
}
