package climbingstairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <= n; i++ {
		current := prev2 + prev1
		prev2, prev1 = prev1, current
	}
	return prev1
}

var cache = make(map[int]int)

func climbStairsRecursion(n int) int {
	if n <= 2 {
		return n
	}

	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = climbStairsRecursion(n-2) + climbStairsRecursion(n-1)
	return cache[n]
}
