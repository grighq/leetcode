package symmetrictree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	root    *TreeNode
	expected bool
}{
	{
		root:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
		expected: true,
	},
	{
		root:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		expected: false,
	},
	{
		root:    &TreeNode{Val: 1},
		expected: true,
	},
	{
		root:    nil,
		expected: true,
	},
	{
		root:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}},
		expected: true,
	},
	{
		root:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		expected: false,
	},
	{
		root:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}},
		expected: false,
	},
}

func TestIsSymmetric(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, isSymmetric(c.root))
	}
}

func TestIsSymmetricRecursion(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, isSymmetricRecursion(c.root))
	}
}