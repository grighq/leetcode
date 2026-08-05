package validboomerang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBoomerang(t *testing.T) {
	assert.Equal(t, true, isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
	assert.Equal(t, false, isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}}))
}
