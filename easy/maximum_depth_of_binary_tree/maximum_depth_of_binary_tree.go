package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		depth++
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

func maxDepthRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepthRecursion(root.Left), maxDepthRecursion(root.Right))
}
