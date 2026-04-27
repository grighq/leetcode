package mergedtwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	list1    []int
	list2    []int
	expected []int
}{
	{[]int{}, []int{}, []int{}},
	{[]int{}, []int{0}, []int{0}},
	{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
}

func TestMergeTwoLists(t *testing.T) {
	for _, c := range cases {
		l1 := sliceToList(c.list1)
		l2 := sliceToList(c.list2)
		output := listToSlice(mergeTwoLists(l1, l2))
		assert.Equal(t, c.expected, output)
	}
}

func TestMergeTwoListsRecursion(t *testing.T) {
	for _, c := range cases {
		l1 := sliceToList(c.list1)
		l2 := sliceToList(c.list2)
		output := listToSlice(mergeTwoListsRecursion(l1, l2))
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
