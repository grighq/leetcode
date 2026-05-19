package movezeroes

func moveZeroes(nums []int) []int {
	k := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	return nums
}
