package longestcontinuousincreasingsubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLengthOfLCIS(t *testing.T) {
	assert.Equal(t, 1, findLengthOfLCIS([]int{2, 2, 2, 2, 2, 2, 2}))
	assert.Equal(t, 3, findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	assert.Equal(t, 4, findLengthOfLCIS([]int{1, 3, 5, 7}))
}
