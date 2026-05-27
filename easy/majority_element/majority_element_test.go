package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	assert.Equal(t, 1, majorityElement([]int{1, 2, 1, 3, 1, 4, 1, 5}))
}
