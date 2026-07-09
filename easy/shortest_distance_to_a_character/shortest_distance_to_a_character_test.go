package shortestdistancetoacharacter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestToChar(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1, 0}, shortestToChar("aaab", 'b'))
	assert.Equal(t, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}, shortestToChar("loveleetcode", 'e'))
}
