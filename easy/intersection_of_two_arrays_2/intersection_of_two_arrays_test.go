package intersectionoftwoarrays2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	assert.Equal(t, []int{2, 2}, intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	assert.Equal(t, []int{4, 9}, intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
