package findthetownjudge

func findJudge(n int, trust [][]int) int {
	trustedScores := make([]int, n+1)
	for _, pair := range trust {
		trustedScores[pair[0]]--
		trustedScores[pair[1]]++
	}

	for i, score := range trustedScores[1:] {
		if score == n-1 {
			return i + 1
		}
	}

	return -1
}
