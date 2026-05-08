package besttimetobuyandsellstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{1}))
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
	assert.Equal(t, 1, maxProfit([]int{1, 2}))
	assert.Equal(t, 0, maxProfit([]int{2, 1}))
	assert.Equal(t, 0, maxProfit([]int{3, 3, 3, 3}))
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
}
