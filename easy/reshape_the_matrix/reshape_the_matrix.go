package reshapethematrix

func matrixReshape(matrix [][]int, rows, cols int) [][]int {
	if len(matrix)*len(matrix[0]) != rows*cols {
		return matrix
	}

	res := make([][]int, rows)
	matRowIdx, matColIdx, matCols := 0, 0, len(matrix[0])

	for i := range rows {
		row := make([]int, cols)
		for idx := range cols {
			if matColIdx >= matCols {
				matRowIdx++
				matColIdx = 0
			}
			row[idx] = matrix[matRowIdx][matColIdx]
			matColIdx++
		}
		res[i] = row
	}

	return res
}

// func matrixReshape(matrix [][]int, rows, cols int) [][]int {
// 	oldRows, oldCols := len(matrix), len(matrix[0])
// 	if oldRows*oldCols != rows*cols {
// 		return matrix
// 	}
//
// 	res := make([][]int, rows)
// 	for i := range res {
// 		res[i] = make([]int, cols)
// 	}
//
// 	totalElements := rows * cols
// 	for i := 0; i < totalElements; i++ {
// 		oldR := i / oldCols
// 		oldC := i % oldCols
//
// 		newR := i / cols
// 		newC := i % cols
//
// 		res[newR][newC] = matrix[oldR][oldC]
// 	}
//
// 	return res
// }
