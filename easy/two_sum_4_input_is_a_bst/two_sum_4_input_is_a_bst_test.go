package twosum4inputisabst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input []int
	k     int
	want  bool
}{
	{[]int{}, 0, false},
	{[]int{1}, 2, false},
	{[]int{5, 3, 6, 2, 4, 0, 7}, 9, true},
	{[]int{5, 3, 6, 2, 4, 0, 7}, 28, false},
	{[]int{2, 1, 3}, 4, true},
	{[]int{2, 1, 3}, 1, false},
	{[]int{0, -1, 1}, 0, true},
	{[]int{1, 0, 2}, 3, true},
	{[]int{3, 1, 4, 0, 2}, 3, true},
	{[]int{3, 1, 4, 0, 2}, 10, false},
}

func TestFindTarget(t *testing.T) {
	for _, c := range cases {
		root := sliceToTree(c.input)
		got := findTarget(root, c.k)
		assert.Equal(t, c.want, got, "input=%v k=%d", c.input, c.k)
	}
}

// sliceToTree builds a binary tree from a level-order slice.
// A value of 0 represents a nil node.
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
