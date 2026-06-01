package findallnumbersdisappearedinanarray

func findDisappearedNumbers(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		i := abs(num) - 1
		nums[i] = min(-nums[i], nums[i])
	}

	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func findDisappearedNumbers(nums []int) []int {
// 	res := make([]int, len(nums))
// 	for _, num := range nums {
// 		res[num-1] = num
// 	}
//
// 	idx := 0
//
// 	for i := range res {
// 		if res[i] == 0 {
// 			res[idx] = i + 1
// 			idx++
// 		}
// 	}
//
// 	return res[:idx]
// }
