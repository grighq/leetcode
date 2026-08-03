package binaryprefixdivisibleby5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixesDivBy5(t *testing.T) {
	assert.Equal(t, []bool{true, false, false}, prefixesDivBy5([]int{0, 1, 1}))
	assert.Equal(t, []bool{false, false, false}, prefixesDivBy5([]int{1, 1, 1}))
}
