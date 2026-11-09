package distancebetweenbusstops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	d1, d2 := 0, 0
	for i, d := range distance {
		if i >= start && i < destination {
			d1 += d
		} else {
			d2 += d
		}
	}

	return min(d1, d2)
}
