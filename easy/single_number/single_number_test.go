package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, 1, singleNumber([]int{1}))
	assert.Equal(t, 1, singleNumber([]int{2, 2, 1}))
	assert.Equal(t, 4, singleNumber([]int{2, 2, 4, 1, 1}))
}
