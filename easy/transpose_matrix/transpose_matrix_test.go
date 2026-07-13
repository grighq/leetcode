package transposematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	assert.Equal(t, [][]int{{1, 4}, {2, 5}, {3, 6}}, transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
	assert.Equal(t, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}, transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
