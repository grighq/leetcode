package toeplitzmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsToeplitzMatrix(t *testing.T) {
	assert.Equal(t, false, isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
	assert.Equal(t, true, isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
}
