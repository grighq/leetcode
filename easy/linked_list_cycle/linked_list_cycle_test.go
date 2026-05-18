package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input []int
	pos   int
	want  bool
}{
	{[]int{}, -1, false},
	{[]int{1}, -1, false},
	{[]int{1}, 0, true},
	{[]int{3, 2, 0, -4}, 1, true},
	{[]int{3, 2, 0, -4}, -1, false},
	{[]int{1, 2}, 0, true},
	{[]int{1, 2}, -1, false},
}

func TestHasCycle(t *testing.T) {
	for _, c := range cases {
		head := sliceToListWithCycle(c.input, c.pos)
		assert.Equal(t, c.want, hasCycle(head))
	}
}

// sliceToListWithCycle creates a linked list from nums and links the
// tail's Next to the node at the given position (0-based). If pos is -1
// or the list is empty, no cycle is created.
func sliceToListWithCycle(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(nums))
	head := &ListNode{Val: nums[0]}
	nodes[0] = head
	curr := head

	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
		nodes[i] = curr
	}

	if pos >= 0 && pos < len(nums) {
		curr.Next = nodes[pos]
	}

	return head
}
