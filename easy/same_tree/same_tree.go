package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Nodes struct {
	node1 *TreeNode
	node2 *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := []Nodes{{p, q}}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		n1, n2 := curr.node1, curr.node2
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}

		stack = append(stack, Nodes{n1.Left, n2.Left})
		stack = append(stack, Nodes{n1.Right, n2.Right})
	}
	return true
}

func isSameTreeRecursion(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTreeRecursion(p.Left, q.Left) && isSameTreeRecursion(p.Right, q.Right)
}
