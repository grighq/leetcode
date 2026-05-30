package rangesumqueryimmutable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSumQueryImmutable(t *testing.T) {
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	assert.Equal(t, 1, n.SumRange(0, 2))
	assert.Equal(t, -1, n.SumRange(2, 5))
	assert.Equal(t, -3, n.SumRange(0, 5))

}
