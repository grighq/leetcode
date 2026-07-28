package availablecapturesforrook

func numRookCaptures(board [][]byte) int {
	for i := range 8 {
		for j := range 8 {
			if board[i][j] == 'R' {
				return countCaptures(board, i, j)
			}
		}
	}

	return 0
}

func countCaptures(board [][]byte, r, c int) int {
	res := 0
	directions := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for _, d := range directions {
		dr, dc := d[0], d[1]
		x, y := r+dr, c+dc

		for x >= 0 && x < 8 && y >= 0 && y < 8 {
			if board[x][y] == 'B' {
				break
			}
			if board[x][y] == 'p' {
				res++
				break
			}

			x += dr
			y += dc
		}
	}

	return res
}
