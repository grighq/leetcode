package removeduplicatesfromsortedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    []int
	expected []int
}{
	{[]int{1, 1, 3}, []int{1, 3}},
	{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
	{[]int{1, 3, 3, 4, 5}, []int{1, 3, 4, 5}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, c := range cases {
		input := sliceToList(c.input)
		output := listToSlice(deleteDuplicates(input))
		assert.Equal(t, c.expected, output)
	}
}

func TestRemoveDuplicatesRecursion(t *testing.T) {
	for _, c := range cases {
		input := sliceToList(c.input)
		output := listToSlice(deleteDuplicatesRecursion(input))
		assert.Equal(t, c.expected, output)
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
