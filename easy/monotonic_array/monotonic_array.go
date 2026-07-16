package monotonicarray

func isMonotonic(nums []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff > 0 {
			decreasing = false
		}

		if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing
}
