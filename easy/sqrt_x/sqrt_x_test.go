package sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 0, mySqrt(0))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 3, mySqrt(9))
	assert.Equal(t, 4, mySqrt(16))
	assert.Equal(t, 4, mySqrt(17))
	assert.Equal(t, 5, mySqrt(30))
	assert.Equal(t, 11, mySqrt(132))
}
