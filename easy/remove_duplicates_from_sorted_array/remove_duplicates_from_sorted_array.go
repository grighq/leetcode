package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	k := 0
	for _, num := range nums[1:] {
		if nums[k] != num {
			k, nums[k+1] = k+1, num
		}
	}
	return k + 1
}
