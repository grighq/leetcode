package heightchecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeightChecker(t *testing.T) {
	assert.Equal(t, 0, heightChecker([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 5, heightChecker([]int{5, 1, 2, 3, 4}))
	assert.Equal(t, 3, heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
