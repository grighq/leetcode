package surfaceareaof3dshapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurfaceArea(t *testing.T) {
	assert.Equal(t, 34, surfaceArea([][]int{{1, 2}, {3, 4}}))
	assert.Equal(t, 32, surfaceArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	assert.Equal(t, 46, surfaceArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}))
}
