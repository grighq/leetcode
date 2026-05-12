package nthtribonaccinumber

func tribonacci(n int) int {
	prev3, prev2, prev1 := 0, 1, 1
	for range n {
		prev3, prev2, prev1 = prev2, prev1, prev3+prev2+prev1
	}
	return prev3
}
