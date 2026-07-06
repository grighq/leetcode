package numberoflinestowritestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfLines(t *testing.T) {
	assert.Equal(t, []int{2, 4}, numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
	assert.Equal(t, []int{3, 60}, numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
}
