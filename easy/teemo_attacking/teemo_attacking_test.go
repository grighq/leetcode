package teemoattacking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPoisonedDuration(t *testing.T) {
	assert.Equal(t, 4, findPoisonedDuration([]int{1, 4}, 2))
	assert.Equal(t, 3, findPoisonedDuration([]int{1, 2}, 2))
}
