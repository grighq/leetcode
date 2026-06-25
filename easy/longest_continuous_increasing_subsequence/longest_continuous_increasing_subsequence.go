package longestcontinuousincreasingsubsequence

func findLengthOfLCIS(nums []int) int {
	maxCIS, currCIS := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currCIS++
		} else {
			currCIS = 1
		}

		maxCIS = max(maxCIS, currCIS)
	}

	return maxCIS
}
