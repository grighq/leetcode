package thirdmaximumnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	assert.Equal(t, 2, thirdMax([]int{1, 2}))
	assert.Equal(t, 2, thirdMax([]int{1, 2, 2}))
	assert.Equal(t, 1, thirdMax([]int{3, 2, 1}))
	assert.Equal(t, 1, thirdMax([]int{2, 2, 3, 1}))
}
