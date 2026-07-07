package mostcommonword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostCommonWord(t *testing.T) {
	assert.Equal(t, "a", mostCommonWord("a.", []string{}))
	assert.Equal(t, "bob", mostCommonWord("bob", []string{}))
	assert.Equal(t, "ball", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}
