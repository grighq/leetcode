package binarytreeinordertraversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input    []int
	expected []int
}{

	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{1, 2, 3}, []int{2, 1, 3}},
	{[]int{1, 2, 3, 4}, []int{4, 2, 1, 3}},
	{[]int{}, []int{}},
}

func TestInorderTraversal(t *testing.T) {
	for _, c := range cases {
		root := sliceToTree(c.input)
		output := inorderTraversal(root)
		assert.Equal(t, c.expected, output)
	}
}

func TestInorderTraversalRecursion(t *testing.T) {
	for _, c := range cases {
		root := sliceToTree(c.input)
		output := inorderTraversalRecursion(root)
		assert.Equal(t, c.expected, output)
	}
}

func sliceToTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != 0 {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != 0 {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

