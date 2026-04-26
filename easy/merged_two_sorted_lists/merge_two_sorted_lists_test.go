package mergedtwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergedTwoLists(t *testing.T) {
	list1 := sliceToList([]int{1, 2, 4})
	list2 := sliceToList([]int{1, 3, 4})
	output := listToSlice(mergeTwoLists(list1, list2))

	list1 = sliceToList([]int{1, 2, 4})
	list2 = sliceToList([]int{1, 3, 4})
	outputRecursion := listToSlice(mergeTwoListsRecursion(list1, list2))

	expected := []int{1, 1, 2, 3, 4, 4}
	assert.Equal(t, expected, output)
	assert.Equal(t, expected, outputRecursion)

	list1 = sliceToList([]int{})
	list2 = sliceToList([]int{})
	output = listToSlice(mergeTwoLists(list1, list2))

	list1 = sliceToList([]int{})
	list2 = sliceToList([]int{})
	outputRecursion = listToSlice(mergeTwoListsRecursion(list1, list2))

	expected = []int{}
	assert.Equal(t, expected, output)
	assert.Equal(t, expected, outputRecursion)

	list1 = sliceToList([]int{})
	list2 = sliceToList([]int{0})
	output = listToSlice(mergeTwoLists(list1, list2))

	list1 = sliceToList([]int{})
	list2 = sliceToList([]int{0})
	outputRecursion = listToSlice(mergeTwoListsRecursion(list1, list2))

	expected = []int{0}
	assert.Equal(t, expected, output)
	assert.Equal(t, expected, outputRecursion)
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
