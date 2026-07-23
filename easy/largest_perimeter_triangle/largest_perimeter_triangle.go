package largestperimetertriangle

import "slices"

func largestPerimeter(nums []int) int {
	slices.Sort(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		leg1, leg2, hypo := nums[i-2], nums[i-1], nums[i]
		if leg1+leg2 > hypo {
			return leg1 + leg2 + hypo
		}
	}

	return 0
}
