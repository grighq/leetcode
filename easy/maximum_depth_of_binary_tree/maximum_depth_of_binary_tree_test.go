package maximumdepthofbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	root     *TreeNode
	expected int
}{
	{
		root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
		expected: 2,
	},
	{
		root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		expected: 3,
	},
	{
		root:     nil,
		expected: 0,
	},
	{
		root:     &TreeNode{Val: 1},
		expected: 1,
	},
	{
		root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}},
		expected: 4,
	},
	{
		root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}},
		expected: 4,
	},
	{
		root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		expected: 3,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, maxDepth(c.root))
	}
}

func TestMaxDepthRecursion(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, maxDepthRecursion(c.root))
	}
}

