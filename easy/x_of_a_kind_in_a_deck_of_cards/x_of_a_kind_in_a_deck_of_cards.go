package xofakindinadeckofcards

func hasGroupsSizeX(deck []int) bool {
	groups := make(map[int]int)
	for _, card := range deck {
		groups[card]++
	}

	minGroup := len(deck)
	for _, group := range groups {
		minGroup = min(minGroup, group)
	}

	for i := minGroup; i > 1; i-- {
		isValid := true
		for _, group := range groups {
			if group%i != 0 {
				isValid = false
				break
			}
		}

		if isValid {
			return true
		}
	}

	return false
}
