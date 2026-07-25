package findthetownjudge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindJudge(t *testing.T) {
	assert.Equal(t, 1, findJudge(1, [][]int{}))
	assert.Equal(t, 2, findJudge(2, [][]int{{1, 2}}))
	assert.Equal(t, 3, findJudge(3, [][]int{{1, 3}, {2, 3}}))
	assert.Equal(t, -1, findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}
