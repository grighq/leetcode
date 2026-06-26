package baseballgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalPoints(t *testing.T) {
	assert.Equal(t, 0, calPoints([]string{"1", "C"}))
	assert.Equal(t, 30, calPoints([]string{"5", "2", "C", "D", "+"}))
	assert.Equal(t, 27, calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
