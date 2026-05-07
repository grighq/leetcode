package pascalstriangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, generate(1))
	assert.Equal(t, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, generate(5))
}
