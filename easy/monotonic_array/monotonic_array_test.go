package monotonicarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMonotonic(t *testing.T) {
	assert.Equal(t, true, isMonotonic([]int{1}))
	assert.Equal(t, false, isMonotonic([]int{1, 3, 2}))
	assert.Equal(t, true, isMonotonic([]int{6, 5, 4, 4}))
	assert.Equal(t, false, isMonotonic([]int{1, 2, 3, 5, 4}))
}
