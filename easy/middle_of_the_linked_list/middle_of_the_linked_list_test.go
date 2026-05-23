package middlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input []int
	want  int
}{
	{[]int{}, 0},
	{[]int{1}, 1},
	{[]int{1, 2}, 2},
	{[]int{1, 2, 3}, 2},
	{[]int{1, 2, 3, 4}, 3},
	{[]int{1, 2, 3, 4, 5}, 3},
}

func TestMiddleNode(t *testing.T) {
	for _, c := range cases {
		head := sliceToList(c.input)

		got := middleNode(head)

		if len(c.input) == 0 {
			assert.Nil(t, got, "input=%v", c.input)
		} else {
			if assert.NotNil(t, got, "input=%v", c.input) {
				assert.Equal(t, c.want, got.Val, "input=%v", c.input)
			}
		}
	}
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head

	for _, num := range nums[1:] {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}

	return head
}
