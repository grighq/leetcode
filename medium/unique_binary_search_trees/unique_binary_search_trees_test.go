package uniquebinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTrees(t *testing.T) {
	assert.Equal(t, 1, numTrees(1))
	assert.Equal(t, 5, numTrees(3))
}
