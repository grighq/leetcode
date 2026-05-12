package nthtribonaccinumber

func tribonacci(n int) int {
	curr, next1, next2 := 0, 1, 1
	for range n {
		curr, next1, next2 = next1, next2, curr+next1+next2
	}

	return curr
}
