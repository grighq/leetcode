package convertbinarynumberinalinkedlisttointeger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input []int
	want  int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{1}, 1},
	{[]int{1, 0}, 2},
	{[]int{1, 0, 1}, 5},
	{[]int{1, 1, 1}, 7},
	{[]int{0, 0}, 0},
	{[]int{1, 0, 1, 0}, 10},
	{[]int{0, 1, 1}, 3},
}

func TestGetDecimalValue(t *testing.T) {
	for _, c := range cases {
		head := sliceToList(c.input)
		got := getDecimalValue(head)
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
