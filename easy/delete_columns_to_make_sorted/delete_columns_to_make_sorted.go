package deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	res := 0
	str := strs[0]

	for i := range str {
		char := str[i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] >= char {
				char = strs[j][i]
			} else {
				res++
				break
			}
		}
	}

	return res
}
