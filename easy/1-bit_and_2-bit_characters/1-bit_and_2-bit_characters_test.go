package bitand2bitcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsOneBitCharacter(t *testing.T) {
	assert.Equal(t, true, isOneBitCharacter([]int{1, 0, 0}))
	assert.Equal(t, false, isOneBitCharacter([]int{1, 1, 1, 0}))
}
