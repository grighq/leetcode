package longestharmonioussubsequence

func findLHS(nums []int) int {
	lhs := 0
	counts := make(map[int]int, len(nums))

	for _, num := range nums {
		counts[num]++
	}

	for num := range counts {
		if count, ok := counts[num+1]; ok {
			lhs = max(lhs, counts[num]+count)
		}
	}

	return lhs
}
