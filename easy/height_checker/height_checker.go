package heightchecker

func heightChecker(heights []int) int {
	counts := [101]int{}
	for _, h := range heights {
		counts[h]++
	}

	res := 0
	currHeight := 1

	for _, h := range heights {
		for counts[currHeight] == 0 {
			currHeight++
		}

		if h != currHeight {
			res++
		}

		counts[currHeight]--
	}

	return res
}
