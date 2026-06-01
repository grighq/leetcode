package thirdmaximumnumber

// Pointers
func thirdMax(nums []int) int {
	var first, second, third *int
	for _, num := range nums {
		if (first != nil && *first == num) || (second != nil && *second == num) || (third != nil && *third == num) {
			continue
		}

		if first == nil || num > *first {
			first, second, third = &num, first, second
		} else if second == nil || num > *second {
			second, third = &num, second
		} else if third == nil || num > *third {
			third = &num
		}
	}

	if third == nil {
		return *first
	}

	return *third
}

// Sorting slices
// func thirdMax(nums []int) int {
// 	slices.Sort(nums)
// 	count, currMax := 1, nums[len(nums)-1]
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		if nums[i] != currMax {
// 			currMax = nums[i]
// 			count++
// 		}
//
// 		if count == 3 {
// 			return currMax
// 		}
// 	}
//
// 	return nums[len(nums)-1]
// }

// Math.MinInt initialization
// func thirdMax(nums []int) int {
// 	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
// 	for _, num := range nums {
// 		if num > first {
// 			first, second, third = num, first, second
// 		} else if num > second && num < first {
// 			second, third = num, second
// 		} else if num > third && num < second {
// 			third = num
// 		}
// 	}
//
// 	if third == math.MinInt64 {
// 		return first
// 	}
//
// 	return third
// }
