package reverselinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{1, 2, 3}, []int{3, 2, 1}},
	{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
}

func TestReverseList(t *testing.T) {
	for _, c := range cases {
		input := sliceToList(c.input)
		output := listToSlice(reverseList(input))
		assert.Equal(t, c.expected, output, "input=%v", c.input)
	}
}

func TestReverseListRecursion(t *testing.T) {
	for _, c := range cases {
		input := sliceToList(c.input)
		output := listToSlice(reverseListRecursion(input))
		assert.Equal(t, c.expected, output, "input=%v", c.input)
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

func listToSlice(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}
