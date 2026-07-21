package distringmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiStringMatch(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3}, diStringMatch("III"))
	assert.Equal(t, []int{3, 2, 0, 1}, diStringMatch("DDI"))
	assert.Equal(t, []int{0, 4, 1, 3, 2}, diStringMatch("IDID"))
}
