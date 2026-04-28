package findindexofthefirstoccurrenceinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndexOfTheFirstOccurenceInAString(t *testing.T) {
	assert.Equal(t, 0, strStr("s", "s"))
	assert.Equal(t, 2, strStr("hell", "ll"))
	assert.Equal(t, 0, strStr("sadbutsac", "sad"))
	assert.Equal(t, -1, strStr("leetcode", "leeto"))
}
