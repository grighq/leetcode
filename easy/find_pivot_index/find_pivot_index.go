package findpivotindex

func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, 0

	for _, num := range nums {
		rightSum += num
	}

	for i := range nums {

		rightSum -= nums[i]

		if leftSum == rightSum {
			return i
		}

		leftSum += nums[i]

	}

	return -1
}
