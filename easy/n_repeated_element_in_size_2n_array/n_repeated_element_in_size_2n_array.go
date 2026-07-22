package nrepeatedelementinsize2narray

func repeatedNTimes(nums []int) int {
	res := 0
	uniqueNums := make(map[int]struct{}, len(nums)/2+1)
	for _, num := range nums {
		if _, ok := uniqueNums[num]; ok {
			res = num
			break
		}
		uniqueNums[num] = struct{}{}
	}

	return res
}

// func repeatedNTimes(nums []int) int {
// 	for i := 2; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] || nums[i] == nums[i-2] {
// 			return nums[i]
// 		}
// 	}
//
// 	return nums[0]
// }
