package arraypartition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	assert.Equal(t, 4, arrayPairSum([]int{1, 4, 3, 2}))
	assert.Equal(t, 9, arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}
