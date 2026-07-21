package validmountainarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidMountainArray(t *testing.T) {
	assert.Equal(t, false, validMountainArray([]int{2, 1}))
	assert.Equal(t, false, validMountainArray([]int{3, 4, 5}))
	assert.Equal(t, true, validMountainArray([]int{0, 5, 3, 2, 1}))
}
