package pascalstriangle2

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex)
	for i := range rowIndex + 1 {
		curr := make([]int, i+1)
		curr[0], curr[i] = 1, 1
		for j := 1; j < i; j++ {
			curr[j] = row[j-1] + row[j]
		}
		row = curr
	}
	return row
}
