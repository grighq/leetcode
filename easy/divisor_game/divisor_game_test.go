package divisorgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisorGame(t *testing.T) {
	assert.Equal(t, true, divisorGame(2))
	assert.Equal(t, false, divisorGame(3))
}
