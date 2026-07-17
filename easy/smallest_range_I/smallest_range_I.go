package smallestrangeI

func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]

	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}

		if num < minNum {
			minNum = num
		}
	}

	diff := maxNum - minNum - 2*k
	if diff <= 0 {
		return 0
	}

	return diff
}
