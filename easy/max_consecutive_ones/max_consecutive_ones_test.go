package maxconsecutiveones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 2, findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
	assert.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
