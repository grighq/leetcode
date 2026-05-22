package removelinkedlistelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    []int
	val      int
	expected []int
}{
	{[]int{}, 1, []int{}},
	{[]int{1}, 1, []int{}},
	{[]int{1}, 2, []int{1}},
	{[]int{1, 2, 3}, 1, []int{2, 3}},
	{[]int{1, 2, 3}, 3, []int{1, 2}},
	{[]int{7, 7, 7, 7}, 7, []int{}},
	{[]int{1, 2, 3, 2, 4}, 2, []int{1, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},
}

func TestRemoveElements(t *testing.T) {
	for _, c := range cases {
		input := sliceToList(c.input)
		output := listToSlice(removeElements(input, c.val))
		assert.Equal(t, c.expected, output, "input=%v val=%d", c.input, c.val)
	}
}

func TestRemoveElementsRecursion(t *testing.T) {
	for _, c := range cases {
		input := sliceToList(c.input)
		output := listToSlice(removeElementsRecursion(input, c.val))
		assert.Equal(t, c.expected, output, "input=%v val=%d", c.input, c.val)
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
