package imagesmoother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageSmoother(t *testing.T) {
	assert.Equal(t, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	assert.Equal(t, [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}}, imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
}
