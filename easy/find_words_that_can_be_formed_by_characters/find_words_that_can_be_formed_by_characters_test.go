package findwordsthatcanbeformedbycharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacters(t *testing.T) {
	assert.Equal(t, 6, countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	assert.Equal(t, 10, countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}
