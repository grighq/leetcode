package largestnumberatleasttwiceofothers

func dominantIndex(nums []int) int {
	maxIdx, secMaxIdx := 0, 1

	for i, num := range nums {
		if num > nums[maxIdx] {
			maxIdx, secMaxIdx = i, maxIdx
		} else if num > nums[secMaxIdx] && num != nums[maxIdx] {
			secMaxIdx = i
		}
	}

	if nums[maxIdx] >= nums[secMaxIdx]*2 {
		return maxIdx
	}

	return -1
}
