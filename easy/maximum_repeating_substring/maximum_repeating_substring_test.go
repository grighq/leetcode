package maximumrepeatingsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxRepeating(t *testing.T) {
	assert.Equal(t, 1, maxRepeating("a", "a"))
	assert.Equal(t, 0, maxRepeating("a", "b"))
	assert.Equal(t, 0, maxRepeating("ababc", "ac"))
	assert.Equal(t, 1, maxRepeating("ababac", "ac"))
	assert.Equal(t, 1, maxRepeating("ababc", "ba"))
	assert.Equal(t, 2, maxRepeating("ababc", "ab"))
	assert.Equal(t, 5, maxRepeating("aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"))
}
