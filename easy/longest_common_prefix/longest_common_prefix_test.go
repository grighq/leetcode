package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "a", longestCommonPrefix([]string{"a"}))
	assert.Equal(t, "a", longestCommonPrefix([]string{"ab", "a"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
