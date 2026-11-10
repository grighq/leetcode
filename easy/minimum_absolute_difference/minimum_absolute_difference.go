package minimumabsolutedifference

import "slices"

func minimumAbsDifference(arr []int) [][]int {
	slices.Sort(arr)
	res := make([][]int, 0, len(arr)/2)
	minDiff := arr[1] - arr[0]
	res = append(res, []int{arr[0], arr[1]})

	for i := 2; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		if diff < minDiff {
			minDiff = diff
			res = res[:0]
			res = append(res, []int{arr[i-1], arr[i]})
		} else if diff == minDiff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}
