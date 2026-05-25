package minimumpairremovaltosortarray1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	assert.Equal(t, 0, minimumPairRemoval([]int{}))
	assert.Equal(t, 0, minimumPairRemoval([]int{15}))
	assert.Equal(t, 0, minimumPairRemoval([]int{1, 2, 2}))
	assert.Equal(t, 2, minimumPairRemoval([]int{5, 2, 3, 1}))
}
