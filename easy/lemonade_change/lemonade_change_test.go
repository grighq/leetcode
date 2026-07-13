package lemonadechange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonadeChange(t *testing.T) {
	assert.Equal(t, true, lemonadeChange([]int{5, 5, 5, 10, 20}))
	assert.Equal(t, false, lemonadeChange([]int{5, 5, 10, 10, 20}))
}
