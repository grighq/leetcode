package twosum4inputisabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, num int) bool {
	nums := []int{}
	var getNums func(*TreeNode)

	getNums = func(r *TreeNode) {
		if r == nil {
			return
		}
		getNums(r.Left)
		nums = append(nums, r.Val)
		getNums(r.Right)
	}

	getNums(root)
	p1, p2 := 0, len(nums)-1
	for p1 < p2 {
		if nums[p1]+nums[p2] == num {
			return true
		}
		if nums[p1]+nums[p2] > num {
			p2--
		} else {
			p1++
		}

	}

	return false
}
