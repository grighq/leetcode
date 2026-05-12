package nthtribonaccinumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	assert.Equal(t, 0, tribonacci(0))
	assert.Equal(t, 2, tribonacci(3))
	assert.Equal(t, 4, tribonacci(4))
	assert.Equal(t, 1389537, tribonacci(25))
}
