package minimumpairremovaltosortarray1

func minimumPairRemoval(nums []int) int {
	count := 0
	for !isSort(nums) {
		count++
		idx := 0
		minSum := nums[0] + nums[1]
		for i := 1; i < len(nums)-1; i++ {
			if nums[i]+nums[i+1] < minSum {
				minSum = nums[i] + nums[i+1]
				idx = i
			}
		}
		nums[idx] = minSum
		nums = append(nums[:idx+1], nums[idx+2:]...)
	}

	return count
}

func isSort(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	prev := nums[0]
	for _, num := range nums[1:] {
		if prev > num {
			return false
		}
		prev = num
	}

	return true
}
