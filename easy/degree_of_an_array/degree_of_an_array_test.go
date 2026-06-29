package degreeofanarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestSubArray(t *testing.T) {
	assert.Equal(t, 2, findShortestSubArray([]int{1, 2, 2, 3, 1}))
	assert.Equal(t, 6, findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
