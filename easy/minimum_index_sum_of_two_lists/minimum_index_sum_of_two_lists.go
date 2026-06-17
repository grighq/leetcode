package minimumindexsumoftwolists

func findRestaurant(list1, list2 []string) []string {
	res := []string{}
	m := make(map[string]int, len(list1))

	for i, s := range list1 {
		m[s] = i
	}

	minSum := len(list1) + len(list2)
	for i, s := range list2 {
		if idx, ok := m[s]; ok {
			currSum := i + idx
			if minSum > currSum {
				res = []string{s}
				minSum = currSum
			} else if minSum == currSum {
				res = append(res, s)
			}
		}
	}

	return res
}
