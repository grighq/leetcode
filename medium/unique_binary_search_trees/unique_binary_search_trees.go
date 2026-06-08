package uniquebinarysearchtree

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for root := 2; root <= n; root++ {
		for i := 1; i <= root; i++ {
			dp[root] += dp[i-1] * dp[root-i]
		}
	}

	return dp[len(dp)-1]
}
