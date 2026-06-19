package maximumproductofthreenumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumProduct(t *testing.T) {
	assert.Equal(t, 6, maximumProduct([]int{1, 2, 3}))
	assert.Equal(t, 24, maximumProduct([]int{1, 2, 3, 4}))
	assert.Equal(t, -6, maximumProduct([]int{-1, -2, -3}))
}
