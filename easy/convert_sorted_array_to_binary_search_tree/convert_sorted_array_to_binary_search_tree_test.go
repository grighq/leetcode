package convertsortedarraytobinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = [][]int{
	{1, 3},
	{-10, -3, 0, 5, 9},
}

func TestSortedArrayToBST(t *testing.T) {
	for _, c := range cases {
		bst := sortedArrayToBST(c)
		assert.True(t, isBalanced(bst))
		assert.Equal(t, c, TreeToSlice(bst))
	}

}

func TreeToSlice(root *TreeNode) []int {
	nums := []int{}
	var toSlice func(*TreeNode)

	toSlice = func(t *TreeNode) {
		if t == nil {
			return
		}

		toSlice(t.Left)
		nums = append(nums, t.Val)
		toSlice(t.Right)
	}

	toSlice(root)
	return nums
}

func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != -1
}

func checkHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := checkHeight(root.Left)
	if left == -1 {
		return -1
	}

	right := checkHeight(root.Right)
	if right == -1 {
		return -1
	}

	diff := max(left-right, right-left)
	if diff > 1 {
		return -1
	}

	return max(left, right) + 1
}
