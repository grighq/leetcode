package toeplitzmatrix

func isToeplitzMatrix(matrix [][]int) bool {
	for row, col := len(matrix)-1, 0; row > 0 || col < len(matrix[0])-1; {
		target := matrix[row][col]
		for r, c := row+1, col+1; r < len(matrix) && c < len(matrix[0]); {
			if target != matrix[r][c] {
				return false
			}
			r++
			c++
		}

		if row > 0 {
			row--
		} else {
			col++
		}

	}

	return true
}
