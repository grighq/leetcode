package imagesmoother

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))

	for rowIndex, row := range img {
		res[rowIndex] = make([]int, len(row))
		for colIndex := range row {

			sum, count := 0, 0
			for i := max(rowIndex-1, 0); i <= min(rowIndex+1, len(img)-1); i++ {
				for j := max(colIndex-1, 0); j <= min(colIndex+1, len(row)-1); j++ {
					count++
					sum += img[i][j]
				}
			}

			res[rowIndex][colIndex] = sum / count
		}
	}

	return res
}
