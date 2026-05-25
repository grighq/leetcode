package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	l1   []int
	l2   []int
	want []int
}{
	{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
	{[]int{0}, []int{0}, []int{0}},
	{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	{[]int{}, []int{1}, []int{1}},
	{[]int{5}, []int{5}, []int{0, 1}},
	{[]int{1, 8}, []int{0}, []int{1, 8}},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, c := range cases {
		l1 := sliceToList(c.l1)
		l2 := sliceToList(c.l2)
		got := listToSlice(addTwoNumbers(l1, l2))
		assert.Equal(t, c.want, got, "l1=%v l2=%v", c.l1, c.l2)
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
