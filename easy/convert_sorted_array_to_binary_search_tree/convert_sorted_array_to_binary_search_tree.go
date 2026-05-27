package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var toBST func(l, r int) *TreeNode

	toBST = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		mid := (l + r) / 2
		root := &TreeNode{}

		root.Left = toBST(l, mid-1)
		root.Val = nums[mid]
		root.Right = toBST(mid+1, r)

		return root
	}

	return toBST(0, len(nums)-1)
}

// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	mid := len(nums) / 2
// 	root := &TreeNode{Val: nums[mid]}
// 	root.Left = sortedArrayToBST(nums[:mid])
// 	root.Right = sortedArrayToBST(nums[mid+1:])
//
// 	return root
// }
