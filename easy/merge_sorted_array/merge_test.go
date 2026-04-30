package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}{
	{[]int{1}, 1, []int{}, 0, []int{1}},
	{[]int{0}, 0, []int{1}, 1, []int{1}},
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
}

func TestMerge(t *testing.T) {
	for _, c := range cases {
		nums1 := c.nums1
		merge(nums1, c.m, c.nums2, c.n)
		assert.Equal(t, c.expected, nums1)
	}
}
