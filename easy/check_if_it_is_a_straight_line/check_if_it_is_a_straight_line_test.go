package checkifitisastraightline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStraightLine(t *testing.T) {
	assert.Equal(t, true, checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}))
	assert.Equal(t, false, checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
}
