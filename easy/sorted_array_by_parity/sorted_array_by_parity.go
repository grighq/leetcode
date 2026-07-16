package sortedarraybyparity

func sortArrayByParity(nums []int) []int {
	p := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

	return nums
}
