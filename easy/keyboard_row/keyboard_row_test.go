package keyboardrow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	assert.Equal(t, []string{}, findWords([]string{"omk"}))
	assert.Equal(t, []string{"adsdf", "sfd"}, findWords([]string{"adsdf", "sfd"}))
	assert.Equal(t, []string{"Alaska", "Dad"}, findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
