package uniquemorsecodewords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueMorseCodeWords(t *testing.T) {
	assert.Equal(t, 1, uniqueMorseRepresentations([]string{"a"}))
	assert.Equal(t, 2, uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
