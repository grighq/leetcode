package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	prev1, prev2 := cost[0], cost[1]
	for _, c := range cost[2:] {
		curr := c + min(prev1, prev2)
		prev1, prev2 = prev2, curr
	}

	return min(prev1, prev2)
}
