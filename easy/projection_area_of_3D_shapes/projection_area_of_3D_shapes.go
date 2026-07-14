package projectionareaof3dshapes

func projectionArea(grid [][]int) int {
	res := 0
	rowMax := 0
	colMaxes := make([]int, len(grid))

	for _, row := range grid {
		for i, col := range row {
			if col > 0 {
				res++
			}

			rowMax = max(rowMax, col)
			colMaxes[i] = max(colMaxes[i], col)
		}

		res += rowMax
		rowMax = 0
	}

	for _, val := range colMaxes {
		res += val
	}

	return res
}
