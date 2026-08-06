package matrixcellsindisctanceorder

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	maxDist := max(rCenter, rows-1-rCenter) + max(cCenter, cols-1-cCenter)
	buckets := make([][][]int, maxDist+1)

	for row := range rows {
		for col := range cols {
			distance := max(row-rCenter, rCenter-row) + max(col-cCenter, cCenter-col)
			buckets[distance] = append(buckets[distance], []int{row, col})
		}
	}

	res := make([][]int, 0, rows*cols)

	for _, bucket := range buckets {
		res = append(res, bucket...)
	}
	return res
}
