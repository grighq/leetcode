package fibonaccinumber

func fib(n int) int {
	curr, next := 0, 1
	for range n {
		curr, next = next, curr+next
	}
	return curr
}

func fibRecurcsion(n int) int {
	if n < 2 {
		return n
	}
	return fibRecurcsion(n-2) + fibRecurcsion(n-1)
}
