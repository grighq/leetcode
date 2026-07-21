package distringmatch

func diStringMatch(s string) []int {
	low, high := 0, len(s)
	res := make([]int, high+1)

	for i := range s {
		if s[i] == 'I' {
			res[i] = low
			low++
		} else {
			res[i] = high
			high--
		}
	}

	res[len(res)-1] = low
	return res
}
