package sametree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	p        *TreeNode
	q        *TreeNode
	expected bool
}{
	{
		p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		expected: true,
	},
	{
		p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		q:        &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		expected: false,
	},
	{
		p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		q:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		expected: false,
	},
	{
		p:        &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		q:        nil,
		expected: false,
	},
	{
		p:        nil,
		q:        &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		expected: false,
	},
	{
		p:        nil,
		q:        nil,
		expected: true,
	},
}

func TestIsSameTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, isSameTree(c.p, c.q))
	}
}

func TestIsSameTreeRecursion(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.expected, isSameTreeRecursion(c.p, c.q))
	}
}

