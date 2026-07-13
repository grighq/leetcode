package transposematrix

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := range matrix[0] {
		row := make([]int, len(matrix))
		for j := range matrix {
			row[j] = matrix[j][i]
		}

		res[i] = row
	}

	return res
}
