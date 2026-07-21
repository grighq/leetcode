package deletecolumnstomakesorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDeletionSize(t *testing.T) {
	assert.Equal(t, 0, minDeletionSize([]string{"a"}))
	assert.Equal(t, 0, minDeletionSize([]string{"a", "b"}))
	assert.Equal(t, 1, minDeletionSize([]string{"cba", "daf", "ghi"}))
	assert.Equal(t, 3, minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
