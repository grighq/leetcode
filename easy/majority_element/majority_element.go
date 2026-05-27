package majorityelement

func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for _, num := range nums[1:] {
		if count == 0 {
			res = num
		}

		if res == num {
			count++
		} else {
			count--
		}
	}

	return res
}

// func majorityElement(nums []int) int {
// 	m := make(map[int]int)
// 	for _, num := range nums {
// 		m[num]++
// 		if count, _ := m[num]; count > len(nums)/2 {
// 			return num
// 		}
// 	}
//
// 	return 0
// }
