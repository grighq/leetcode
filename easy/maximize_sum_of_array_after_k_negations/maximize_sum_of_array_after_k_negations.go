package maximizesumofarrayafterknegations

import "slices"

func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)

	for i, num := range nums {
		if num < 0 {
			nums[i] = -num
			k--
		}

		if k == 0 {
			break
		}
	}

	if k%2 != 0 {
		minIdx := 0
		for i, num := range nums[1:] {
			if num < nums[minIdx] {
				minIdx = i + 1
			}
		}

		nums[minIdx] = -nums[minIdx]
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}
