package pascalstriangle

func generate(n int) [][]int {
	pt := make([][]int, n)

	for i := range n {
		curr := make([]int, i+1)
		curr[0], curr[i] = 1, 1
		for j := 1; j < i; j++ {
			curr[j] = pt[i-1][j-1] + pt[i-1][j]
		}
		pt[i] = curr
	}

	return pt
}
