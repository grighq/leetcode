package palindromelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input []int
	want  bool
}{
	{[]int{1}, true},
	{[]int{1, 2, 2, 1}, true},
	{[]int{1, 2, 3, 2, 1}, true},
	{[]int{1, 2, 3, 4, 5}, false},
	{[]int{1, 2}, false},
	{[]int{1, 1, 1}, true},
	{[]int{1, 2, 3}, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, c := range cases {
		head := sliceToList(c.input)
		got := isPalindrome(head)
		assert.Equal(t, c.want, got, "input=%v", c.input)
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
