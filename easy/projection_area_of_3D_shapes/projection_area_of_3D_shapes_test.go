package projectionareaof3dshapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectionArea(t *testing.T) {
	assert.Equal(t, 5, projectionArea([][]int{{2}}))
	assert.Equal(t, 8, projectionArea([][]int{{1, 0}, {0, 2}}))
	assert.Equal(t, 17, projectionArea([][]int{{1, 2}, {3, 4}}))
}
