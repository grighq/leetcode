package addtoarrayformofinteger

import (
	"slices"
)

func addToArrayForm(num []int, k int) []int {
	res := make([]int, max(len(num), getLen(k))+1)
	idx := len(res) - 1

	for i := range slices.Backward(num) {
		k += num[i]
		res[idx] = k % 10
		k /= 10
		idx--
	}

	for ; k > 0; k /= 10 {
		res[idx] = k % 10
		idx--
	}

	return res[idx+1:]
}

func getLen(num int) int {
	if num == 0 {
		return 1
	}

	count := 0
	for num > 0 {
		num /= 10
		count++
	}

	return count
}
