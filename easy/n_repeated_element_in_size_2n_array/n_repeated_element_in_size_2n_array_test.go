package nrepeatedelementinsize2narray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedNTimes(t *testing.T) {
	assert.Equal(t, 3, repeatedNTimes([]int{1, 2, 3, 3}))
	assert.Equal(t, 2, repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	assert.Equal(t, 5, repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
}
