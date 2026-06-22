package setmissmatch

func findErrorNums(nums []int) []int {
	m := make(map[int]struct{}, len(nums))
	duplicate, missing := 0, 0

	for _, num := range nums {
		if _, ok := m[num]; ok {
			duplicate = num
		}

		m[num] = struct{}{}
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			missing = i
			break
		}
	}

	return []int{duplicate, missing}
}

// func findErrorNums(nums []int) []int {
// 	actualSum, uniqueSum := 0, 0
// 	neddedSum := len(nums) * (len(nums) + 1) / 2
//
// 	uniqueNums := make(map[int]struct{}, len(nums))
// 	for _, num := range nums {
// 		actualSum += num
// 		if _, ok := uniqueNums[num]; !ok {
// 			uniqueSum += num
// 			uniqueNums[num] = struct{}{}
// 		}
// 	}
//
// 	return []int{actualSum - uniqueSum, neddedSum - uniqueSum}
// }
