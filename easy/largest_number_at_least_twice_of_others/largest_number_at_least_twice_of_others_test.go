package largestnumberatleasttwiceofothers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDominantIndex(t *testing.T) {
	assert.Equal(t, 0, dominantIndex([]int{1, 0}))
	assert.Equal(t, 1, dominantIndex([]int{3, 6, 1, 0}))
	assert.Equal(t, -1, dominantIndex([]int{1, 2, 3, 4}))
	assert.Equal(t, -1, dominantIndex([]int{0, 0, 3, 2}))
}
