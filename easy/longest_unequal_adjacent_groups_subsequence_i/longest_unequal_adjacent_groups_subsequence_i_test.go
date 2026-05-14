package longestunequaladjacentgroupssubsequencei

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLongestSubsequence(t *testing.T) {
	assert.Equal(t, []string{"e"}, getLongestSubsequence([]string{"e"}, []int{0}))
	assert.Equal(t, []string{"e", "b"}, getLongestSubsequence([]string{"e", "a", "b"}, []int{0, 0, 1}))
	assert.Equal(t, []string{"a", "b", "c"}, getLongestSubsequence([]string{"a", "b", "c", "d"}, []int{1, 0, 1, 1}))
}
