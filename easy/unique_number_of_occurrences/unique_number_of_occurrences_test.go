package uniquenumberofoccurrences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	assert.Equal(t, false, uniqueOccurrences([]int{1, 2}))
	assert.Equal(t, true, uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	assert.Equal(t, true, uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
