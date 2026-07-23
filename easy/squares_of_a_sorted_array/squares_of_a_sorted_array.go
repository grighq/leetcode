package squaresofasortedarray

import "slices"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	low, high := 0, len(nums)-1
	for i := range slices.Backward(res) {
		lowSq := nums[low] * nums[low]
		highSq := nums[high] * nums[high]

		if lowSq > highSq {
			res[i] = lowSq
			low++
		} else {
			res[i] = highSq
			high--
		}
	}

	return res
}
