package rangeaddition2

func maxCount(m, n int, ops [][]int) int {
	for _, o := range ops {
		x, y := o[0], o[1]
		m = min(m, x)
		n = min(n, y)
	}

	return m * n
}
