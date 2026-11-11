package checkifitisastraightline

func checkStraightLine(coordinates [][]int) bool {
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[len(coordinates)-1][0], coordinates[len(coordinates)-1][1]
	dx, dy := x2-x1, y2-y1
	for _, point := range coordinates[1 : len(coordinates)-1] {
		x, y := point[0], point[1]
		if dx*(y-y1) != dy*(x-x1) {
			return false
		}
	}

	return true
}
