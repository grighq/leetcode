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
