package containsduplicate2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	assert.True(t, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	assert.True(t, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	assert.False(t, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

}
