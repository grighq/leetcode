package numberofequivalentdominopairs

func numEquivDominoPairs(dominoes [][]int) int {
	res := 0
	counts := [100]int{}

	for _, dom := range dominoes {
		a, b := dom[0], dom[1]
		if a > b {
			a, b = b, a
		}

		key := a*10 + b
		res += counts[key]
		counts[key]++
	}

	return res
}
