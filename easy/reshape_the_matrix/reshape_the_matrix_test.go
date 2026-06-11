package reshapethematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixReshape(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	assert.Equal(t, [][]int{{1, 2}, {3, 4}}, matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4))
}
