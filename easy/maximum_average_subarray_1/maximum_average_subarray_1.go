package maximumaveragesubarray1

func findMaxAverage(nums []int, k int) float64 {
	curSum := 0
	for i := range k {
		curSum += nums[i]
	}

	maxSum := curSum
	for i := k; i < len(nums); i++ {
		curSum = curSum - nums[i-k] + nums[i]
		maxSum = max(maxSum, curSum)
	}

	return float64(maxSum) / float64(k)
}
