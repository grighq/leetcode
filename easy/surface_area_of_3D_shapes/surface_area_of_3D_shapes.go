package surfaceareaof3dshapes

func surfaceArea(grid [][]int) int {
	area := 0
	for i, row := range grid {
		for j, col := range row {
			if col != 0 {
				area += 4*col + 2
			}

			if j > 0 {
				area -= 2 * min(col, row[j-1])
			}

			if i > 0 {
				area -= 2 * min(col, grid[i-1][j])
			}
		}
	}

	return area
}
