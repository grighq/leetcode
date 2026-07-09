package flippinganimage

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		low, high := 0, len(row)-1
		for low <= high {
			row[low], row[high] = 1-row[high], 1-row[low]
			low++
			high--
		}
	}

	return image
}
