package distancebetweenbusstops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceBetweenBusStops(t *testing.T) {
	assert.Equal(t, 1, distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
	assert.Equal(t, 3, distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	assert.Equal(t, 4, distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
}
