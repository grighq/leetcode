package shortestdistancetoacharacter

func shortestToChar(s string, c byte) []int {
	pos := -len(s)
	res := make([]int, len(s))

	for i := range len(s) {
		if s[i] == c {
			pos = i
		}
		res[i] = i - pos
	}

	pos = 2 * len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			pos = i
		}
		res[i] = min(res[i], pos-i)
	}

	return res
}
