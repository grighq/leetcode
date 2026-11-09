package relativesortarray

func relativeSortArray(arr1, arr2 []int) []int {
	counts := [1001]int{}
	for _, num := range arr1 {
		counts[num]++
	}

	res := make([]int, 0, len(arr1))

	for _, num := range arr2 {
		for counts[num] > 0 {
			res = append(res, num)
			counts[num]--
		}
	}

	for i, num := range counts {
		for range num {
			res = append(res, i)
		}
	}

	return res
}
