package maxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	count, currMax := 0, 0
	for _, num := range nums {
		if num == 1 {
			count++
			currMax = max(currMax, count)
		} else {
			count = 0
		}
	}

	return currMax
}
