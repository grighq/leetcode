package minimumcosttomovechipstothesameposition

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, p := range position {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return min(odd, even)
}
