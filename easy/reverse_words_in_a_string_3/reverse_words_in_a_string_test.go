package reversewordsinastring3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "rM gniD", reverseWords("Mr Ding"))
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
}
