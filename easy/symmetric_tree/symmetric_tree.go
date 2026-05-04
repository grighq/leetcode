package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pairs struct {
	left  *TreeNode
	right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []Pairs{{root.Left, root.Right}}
	for len(stack) > 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		left, right := pair.left, pair.right
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		stack = append(stack, Pairs{pair.left.Left, pair.right.Right})
		stack = append(stack, Pairs{pair.left.Right, pair.right.Left})
	}
	return true
}

func isSymmetricRecursion(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isMirror func(left, right *TreeNode) bool

	isMirror = func(left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}

		return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}
